import { HLS_SERVER_URL } from "./config";
import { startStream, stopStream } from "./streamControl";
import "./style.css";
import Hls from "hls.js";

const connectButton = <HTMLButtonElement>(
  document.querySelector("#connect-button")
);
const stopStreamButton = <HTMLButtonElement>(
  document.querySelector("#stop-stream-button")
);
const startStreamButton = <HTMLButtonElement>(
  document.querySelector("#start-stream-button")
);
const statusElement = <HTMLElement>document.querySelector("#status");
const videoElement = <HTMLVideoElement>document.querySelector("#video-stream");
const hls = new Hls();

isStreamOnline().then((isOnline) => {
  if (isOnline) {
    setStreamStatus("Online");
  } else {
    setStreamStatus("Offline");
  }
});

setInterval(() => {
  isStreamOnline().then((isOnline) => {
    if (isOnline) {
      setStreamStatus("Online");
    } else {
      setStreamStatus("Offline");
    }
  });
}, 5000);

let connected = false;

connectButton.addEventListener("click", () => {
  if (connected) {
    connectButton.innerHTML = "Connect";
    disconnectFromVideoStream();
  } else {
    connectButton.innerHTML = "Disconnect";
    connectToVideoStream();
  }

  connected = !connected;
});

stopStreamButton.addEventListener("click", async () => {
  const success = await stopStream();
  if (success) {
    setStreamStatus("Offline");
    if (disconnectFromVideoStream()) {
      connected = false
    }
    connectButton.innerHTML = "Connect";
  }
});

startStreamButton.addEventListener("click", async () => {
  const success = await startStream();

  if (success) {
    setStreamStatus("Online");
  }
});

function connectToVideoStream(): boolean {
  if (!Hls.isSupported()) return false;

  hls.on(Hls.Events.MEDIA_ATTACHED, () => {
    console.log("video and hls.js are now bound together");
  });
  hls.on(Hls.Events.MANIFEST_PARSED, function (_, data) {
    console.log(
      "manifest loaded, found " + data.levels.length + " quality level"
    );
    videoElement.currentTime = hls.liveSyncPosition || 0
    videoElement.play()
  })
  hls.on(Hls.Events.ERROR, function (_, data) {
    console.error(data.type);
    console.error(data.details);
    console.error(data.error);
  });

  hls.loadSource(`${HLS_SERVER_URL}/stream/livestream.m3u8`);
  hls.attachMedia(videoElement);

  return true;
}

function disconnectFromVideoStream(): boolean {
  if (!connected) return false;

  hls.detachMedia();
  hls.stopLoad()
  videoElement.srcObject = null;

  return true;
}

function setStreamStatus(status: string) {
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

async function isStreamOnline(): Promise<boolean> {
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
