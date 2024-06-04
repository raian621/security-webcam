package main

import (
	"flag"
	"fmt"
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
	ffmpegCmd    *exec.Cmd
	stopStream   = make(chan struct{}, 1)
	siteDistPath string
)

func main() {
	host := flag.String("host", "0.0.0.0", "hostname to host the HLS server on")
	port := flag.String("port", "8080", "port to listen on")
	distPath := flag.String("dist-path", "dist", "path to public site files")
	videoDevice := flag.String("video-device", "/dev/video0", "video device to stream")

	flag.Parse()

	siteDistPath = *distPath

	ffmpegArgsInit(*videoDevice)
	createStreamDir()
	go startStream(stopStream)
	streamOnline = true

	server := http.Server{
		Addr:    net.JoinHostPort(*host, *port),
		Handler: http.NewServeMux(),
	}
	addHandlers(server.Handler.(*http.ServeMux))

	log.Printf("Listening on %s://%s:%s", "http", *host, *port)
	log.Fatalln(server.ListenAndServe())

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

	ffmpegCmd = cmd
	log.Printf("started ffmpeg process with pid `%d`\n", cmd.Process.Pid)
	<-stopStream

	log.Printf("stopping ffmpeg process with pid `%d`\n", cmd.Process.Pid)
	if err := cmd.Process.Kill(); err != nil {
		log.Fatalln(err)
	}
	ffmpegCmd = nil

	streamOnline = false
}

func ffmpegArgsInit(videoDevice string) {
	getResolutions(videoDevice)

	ffmpegArgs = []string{
		"-i",
		videoDevice,
		// "-r", "5",
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
			// fmt.Sprintf("scale=w=%d:h=%d:force_original_aspect_ratio=decrease", res.width, res.height),
			fmt.Sprintf("scale=w=%d:h=%d", res.Width, res.Height),
			fmt.Sprintf("-maxrate:v:%d", i),
			fmt.Sprintf("%dk", res.MaxKbps),
		)
	}
	ffmpegArgs = append(ffmpegArgs, "-var_stream_map")
	ffmpegArgs = append(
		ffmpegArgs,
		func() string {
			var varStreamMap strings.Builder

			for i, res := range resolutions {
				varStreamMap.WriteString(
					fmt.Sprintf("v:%d,name:%dp", i, res.Height),
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
