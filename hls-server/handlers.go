package main

import (
	"encoding/json"
	"net/http"
)

func addHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", GET(CORS(logging(
		http.FileServer(http.Dir(siteDistPath)).ServeHTTP,
	))))
	mux.HandleFunc("/stream/", GET(CORS(logging(
		http.StripPrefix(
			"/stream/",
			http.FileServer(http.Dir(streamDir)),
		).ServeHTTP,
	))))
	mux.HandleFunc("/start-stream/", POST(CORS(logging(handleStartStream))))
	mux.HandleFunc("/stop-stream/", POST(CORS(logging(handleStopStream))))
	mux.HandleFunc("/status/", GET(CORS(logging(handleStatus))))
}

func handleStartStream(w http.ResponseWriter, r *http.Request) {
	if !streamOnline {
		go startStream(stopStream)
		streamOnline = true
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func handleStopStream(w http.ResponseWriter, r *http.Request) {
	if streamOnline {
		stopStream <- struct{}{}
		streamOnline = false
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// this should probably be done through websockets, but whatever\
func handleStatus(w http.ResponseWriter, r *http.Request) {
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
}
