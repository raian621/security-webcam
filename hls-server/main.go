package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	streamDir    = "stream"
	streamOnline = false
	ffmpegArgs   = make([]string, 0)
)

func main() {
	host := flag.String("host", "0.0.0.0", "hostname to host the HLS server on")
	port := flag.String("port", "8080", "port to listen on")
	distPath := flag.String("dist-path", "dist", "path to public site files")
	videoDevice := flag.String("video-device", "/dev/video0", "video device to stream")

	flag.Parse()
	log.Println(*distPath)
	stopStream := make(chan struct{}, 1)

	ffmpegArgsInit(*videoDevice)
	createStreamDir()
	go startStream(stopStream)
	streamOnline = true

	http.HandleFunc("/", logging(GET(http.FileServer(http.Dir(*distPath)).ServeHTTP)))

	http.HandleFunc("/stream/", logging(CORS(
		http.StripPrefix(
			"/stream/",
			http.FileServer(http.Dir(streamDir)),
		).ServeHTTP,
	)))

	http.HandleFunc("/start-stream/", logging(CORS(POST(func(w http.ResponseWriter, r *http.Request) {
		if !streamOnline {
			go startStream(stopStream)
			streamOnline = true
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}))))

	http.HandleFunc("/stop-stream/", logging(CORS(POST(func(w http.ResponseWriter, r *http.Request) {
		if streamOnline {
			stopStream <- struct{}{}
			streamOnline = false
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}))))

	// this should probably be done through websockets, but whatever
	http.HandleFunc("/status/", logging(CORS(GET(func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			StreamOnline bool `json:"streamOnline"`
		}{
			StreamOnline: streamOnline,
		}

		w.Header().Add("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(&status)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}))))

	log.Printf("Listening on %s://%s:%s", "http", *host, *port)
	log.Fatalln(http.ListenAndServe(net.JoinHostPort(*host, *port), nil))

	stopStream <- struct{}{}
}

func createStreamDir() {
	os.Mkdir(streamDir, 0744)
}

func startStream(stopStream chan struct{}) {
	cmd := exec.Command("ffmpeg", ffmpegArgs...)

	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("started ffmpeg process with pid `%d`\n", cmd.Process.Pid)
	<-stopStream

	log.Printf("stopping ffmpeg process with pid `%d`\n", cmd.Process.Pid)
	if err := cmd.Process.Kill(); err != nil {
		log.Fatalln(err)
	}

	streamOnline = false
}

func ffmpegArgsInit(videoDevice string) {
	resolutions := []struct {
		width   int
		height  int
		maxKbps int
	}{
		{width: 480, height: 360, maxKbps: 600},
		{width: 640, height: 480, maxKbps: 1500},
		{width: 1280, height: 720, maxKbps: 3000},
		// {width: 1920, height: 1080, maxKbps: 6000},
	}

	ffmpegArgs = []string{
		"-i",
		videoDevice,
		"-c:v",
		"libx264",
		"-crf",
		"22",
		"-c:a",
		"aac",
		"-preset", "fast", "-flags", "+cgop", "-g", "1", "-tune", "zerolatency",
		"-pix_fmt", "yuv420p",
	}

	for range resolutions {
		ffmpegArgs = append(ffmpegArgs, "-map", "0:v:0")
	}

	// add arguments for each output resolution
	for i, res := range resolutions {
		ffmpegArgs = append(
			ffmpegArgs,
			fmt.Sprintf("-filter:v:%d", i),
			fmt.Sprintf("scale=w=%d:h=%d:force_original_aspect_ratio=decrease", res.width, res.height),
			fmt.Sprintf("-maxrate:v:%d", i),
			fmt.Sprintf("%dk", res.maxKbps),
		)
	}
	ffmpegArgs = append(ffmpegArgs, "-var_stream_map")
	ffmpegArgs = append(
		ffmpegArgs,
		func() string {
			var varStreamMap strings.Builder

			for i, res := range resolutions {
				varStreamMap.WriteString(
					fmt.Sprintf("v:%d,name:%dp", i, res.height),
				)
				if i < len(resolutions)-1 {
					varStreamMap.WriteByte(' ')
				}
			}

			return varStreamMap.String()
		}(),
	)

	ffmpegArgs = append(
		ffmpegArgs,
		"-f", "hls",
		"-threads", "0",
		"-hls_list_size", "2",
		"-hls_time", "1",
		"-hls_flags", "delete_segments+independent_segments",
		"-master_pl_name", `livestream.m3u8`,
		"-y", "stream/livestream-%v.ts",
	)
}

func insertHlsUrlIntoBundle(w http.ResponseWriter, filepath, dir, url string) {
	tmpl, err := template.New(filepath).ParseFiles(fmt.Sprintf("%s/%s", dir, filepath))

	if err != nil {
		log.Println(err)
	}

	err = tmpl.Execute(w, struct {
		HlsServerUrl string
	}{HlsServerUrl: url})

	tmpl.Execute(w, struct {
		HlsServerUrl string
	}{HlsServerUrl: url})

	if err != nil {
		log.Println(err)
	}
}
