import {
  isStreamOnline,
  setStreamStatus as _setStreamStatus,
  startStream,
  stopStream,
  disconnectFromVideoStream as _disconnectFromVideoStream,
  connectToVideoStream,
} from "./streamUtils";
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
let connected = false;

const setStreamStatus = (status: string) => {
  return _setStreamStatus(
    status,
    statusElement,
    startStreamButton,
    stopStreamButton,
    connectButton
  );
};
const disconnectFromVideoStream = (): boolean => {
  return _disconnectFromVideoStream(hls, videoElement, connected);
};

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

connectButton.addEventListener("click", () => {
  if (connected) {
    connectButton.innerHTML = "Connect";
    disconnectFromVideoStream();
  } else {
    connectButton.innerHTML = "Disconnect";
    connectToVideoStream(hls, videoElement);
  }

  connected = !connected;
});

stopStreamButton.addEventListener("click", async () => {
  const success = await stopStream();
  if (success) {
    setStreamStatus("Offline");
    if (disconnectFromVideoStream()) {
      connected = false;
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
