# Security Webcam

## Overview

This application allows you to stream your webcam video to any web browser on
your LAN (and probably over the internet). I created this because I occasionally
need to watch for packages or people to arrive at my house but don't really want
to set up a permanent camera or keep looking out of or running over to the
window every 5 minutes.

NOTE: The server-side of this application only works on Linux.

### Client

The browser client for this application is a basic website written in TypeScript,
HTML, and CSS. [Vite](https://vitejs.dev/) is used as a local development server
and asset bundling. It uses the [hls.js](https://github.com/video-dev/hls.js)
library to handle playing and buffering the HLS video streams broadcasted by the
server. 

### Server

The server for this application is written in Go and uses [HTTP Live Streaming (HLS)](https://en.wikipedia.org/wiki/HTTP_Live_Streaming) to broadcast a host computer's webcam.
This is achieved by using [ffmpeg](https://ffmpeg.org/) to transcode video from
the host's webcam into `.m3u8` playlist files and `.ts` segment files, which are
then served to the client via the server's `/stream/` HTTP endpoint.

## Building

To build the application into the `build` subdirectory, run the supplied build script:

```sh
./script/build <HLS server URL>
```

## Development

### Client

To get hot reloads in the browser whilde developing the frontend code, run:

```sh
npm run dev
```

Additionally, you can set the HLS server URL by setting the
`VITE_HLS_SERVER_URL` environment variable to a whatever URL the HLS server is
accessible from (e.g. `http://10.0.0.1:8080` for a LAN URL, and
`https://mydomain.com` for a potential WAN URL).

### Server

To run the backend server with reloads on code changes, you can either run the
provided script:

```sh
./scripts/hls-server-dev
```

or navigate to the `hls-server` subdirectory and running [air](https://github.com/cosmtrek/air)
manually:

```sh
cd hls-server
air
```

You can also just run the server without reloads on code changes by either running

```sh
./scripts/start-hls-server
```

or navigating to the `hls-server` subdirectory and running the server with `go run`:

```sh
cd hls-server
go run .
```

## TODO

- [x] Implement a mechanism to check the max webcam resolution, frame rate, etc.
- [ ] Control webcam resolution, frame rate, etc.
- [ ] Dockerize the application, publish on ghcr.io
