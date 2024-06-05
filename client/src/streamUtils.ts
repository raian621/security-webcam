import Hls from "hls.js";
import { HLS_SERVER_URL } from "./config";

export async function startStream(): Promise<boolean> {
  try {
    const response = await fetch(`${HLS_SERVER_URL}/start-stream/`, {
      method: "POST",
    });
    return response.ok;
  } catch (e) {
    console.error(e);
    return false;
  }
}

export async function stopStream(): Promise<boolean> {
  try {
    const response = await fetch(`${HLS_SERVER_URL}/stop-stream/`, {
      method: "POST",
    });
    return response.ok;
  } catch (e) {
    console.error(e);
    return false;
  }
}

export function connectToVideoStream(
  hls: Hls,
  videoElement: HTMLVideoElement
): boolean {
  if (!Hls.isSupported()) return false;

  hls.on(Hls.Events.MEDIA_ATTACHED, () => {
    console.log("video and hls.js are now bound together");
  });
  hls.on(Hls.Events.MANIFEST_PARSED, function (_, data) {
    console.log(
      "manifest loaded, found " + data.levels.length + " quality level"
    );
    videoElement.currentTime = hls.liveSyncPosition || 0;
    videoElement.play();
  });
  hls.on(Hls.Events.ERROR, function (_, data) {
    console.error(data.type);
    console.error(data.details);
    console.error(data.error);
  });

  hls.loadSource(`${HLS_SERVER_URL}/stream/livestream.m3u8`);
  hls.attachMedia(videoElement);

  return true;
}

export function disconnectFromVideoStream(
  hls: Hls,
  videoElement: HTMLVideoElement,
  connected: boolean,
): boolean {
  if (!connected) return false;

  hls.detachMedia();
  hls.stopLoad();
  videoElement.srcObject = null;

  return true;
}

export function setStreamStatus(
  status: string,
  statusElement: HTMLElement,
  startStreamButton: HTMLButtonElement,
  stopStreamButton: HTMLButtonElement,
  connectButton: HTMLButtonElement
) {
  if (status === "Online") {
    statusElement.className = "online-status";
    startStreamButton.disabled = true;
    stopStreamButton.disabled = false;
    connectButton.disabled = false;
  } else if (status == "Offline") {
    statusElement.className = "offline-status";
    startStreamButton.disabled = false;
    stopStreamButton.disabled = true;
    connectButton.disabled = true;
  }

  statusElement.innerText = status;
}

export async function isStreamOnline(): Promise<boolean> {
  try {
    const res = await fetch(`${HLS_SERVER_URL}/status/`, {
      method: "GET",
    });
    const data = await res.json();
    const { streamOnline }: { streamOnline: boolean } = data;
    return streamOnline;
  } catch (e) {
    console.log(e);
    return false;
  }
}
