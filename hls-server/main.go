package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
)

var (
	streamDir    = "stream"
	streamOnline = false
)

func main() {
	host := flag.String("host", "0.0.0.0", "hostname to host the hls server on")
	port := flag.String("port", "8080", "port to listen on")
	distPath := flag.String("dist-path", "dist", "path to public site files")

	flag.Parse()

	stopStream := make(chan struct{})

	createStreamDir()
	go startStream(stopStream)
	streamOnline = true

	http.HandleFunc("/", logging(GET(
		http.FileServer(http.Dir(*distPath)).ServeHTTP,
	)))

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
	cmd := exec.Command(
		"ffmpeg", "-f", "v4l2",
		"-i", "/dev/video0",
		"-c:v", "libx264",
		"-pix_fmt", "yuv420p",
		"-f", "hls",
		"-hls_time", "5",
		"-hls_playlist_type", "event",
		"stream/output.m3u8",
	)

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

func CORS(nextFn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		nextFn(w, r)
	}
}

func POST(nextFn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		nextFn(w, r)
	}
}

func GET(nextFn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		nextFn(w, r)
	}
}

func logging(nextFn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// timeReceived := time.Now()

		nextFn(w, r)

		method := r.Method
		path := r.URL.Path
		ip := r.RemoteAddr

		log.Printf("%s %s %s\n", method, path, ip)
	}
}
