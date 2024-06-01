import { HLS_SERVER_URL } from "./config";
import { startStream, stopStream } from "./streamControl";
import "./style.css";
import Hls from "hls.js";

const connectButton = <HTMLButtonElement>document.querySelector("#connect-button");
const stopStreamButton = <HTMLButtonElement>document.querySelector("#stop-stream-button");
const startStreamButton = <HTMLButtonElement>document.querySelector("#start-stream-button");
const statusElement = <HTMLElement>document.querySelector("#status");

const videoElement = <HTMLVideoElement>document.querySelector("#video-stream");
const hls = new Hls()

isStreamOnline().then(isOnline => {
  if (isOnline) {
    setStreamStatus("Online")
  } else {
    setStreamStatus("Offline")
  }
})

setInterval(() => {
  isStreamOnline().then(isOnline => {
    if (isOnline) {
      setStreamStatus("Online")
    } else {
      setStreamStatus("Offline")
    }
  })
}, 5000)

let connected = false;

if (connectButton) {
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
}

if (stopStreamButton) {
  stopStreamButton.addEventListener("click", async() => {
    const success = await stopStream()
    if (success) {
      setStreamStatus("Offline")
    }
  })
}

if (startStreamButton) {
  startStreamButton.addEventListener("click", async() => {
    const success = await startStream()

    if (success) {
      setStreamStatus("Online")
    }
  })
}

function connectToVideoStream(): boolean {
  if (!videoElement || !Hls.isSupported()) return false;

  hls.on(Hls.Events.MEDIA_ATTACHED, () => {
    console.log("video and hls.js are now bound together");
  });
  hls.on(Hls.Events.MANIFEST_PARSED, function (_, data) {
    console.log(
      "manifest loaded, found " + data.levels.length + " quality level"
    );
    // (videoElement as HTMLVideoElement).play();
  });
  hls.on(Hls.Events.ERROR, function (_, data) {
    // switch (data.type) {
    //   case Hls.ErrorTypes.KEY_SYSTEM_ERROR:
    //     console.error("HLS key system error occurred")
    //     break;
    //   case Hls.ErrorTypes.MEDIA_ERROR:
    //     console.error("HLS nedia error occurred")
    //     break;
    //   case Hls.ErrorTypes.MUX_ERROR:
    //     console.error("HLS mux error occurred")
    //     break;
    //   case Hls.ErrorTypes.NETWORK_ERROR:
    //     console.error("HLS network error occurred")
    //     break;
    //   case Hls.ErrorTypes.OTHER_ERROR:
    //     console.error("HLS other error occurred")
    //     break;
    // }
    console.error(data.type)
    console.error(data.details)
    console.error(data.error)
  });

  hls.loadSource(`${HLS_SERVER_URL}/stream/output.m3u8`);
  hls.attachMedia(videoElement);

  return true;
}

function disconnectFromVideoStream(): boolean {
  if (!videoElement) return false;

  hls.detachMedia()
  videoElement.srcObject = null

  return true;
}

function setStreamStatus(status: string) {
  if (status === "Online") {
    statusElement.className = "online-status"
    startStreamButton.disabled = true
    stopStreamButton.disabled = false
    connectButton.disabled = false
  } else if (status == "Offline") {
    statusElement.className = "offline-status"
    startStreamButton.disabled = false
    stopStreamButton.disabled = true
    connectButton.disabled = true
  }

  statusElement.innerText = status
}

async function isStreamOnline() : Promise<boolean> {
  try {
    const res = await fetch(`${HLS_SERVER_URL}/status/`, {
      method: "GET",
    })
    const data = await res.json()
    const { streamOnline } : { streamOnline: boolean } = data;
    return streamOnline
  } catch (e) {
    console.log(e)
    return false
  }
}