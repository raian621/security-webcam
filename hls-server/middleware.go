package main

import (
	"log"
	"net/http"
)

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
		nextFn(w, r)

		method := r.Method
		path := r.URL.Path
		ip := r.RemoteAddr

		log.Printf("%s %s %s\n", method, path, ip)
	}
}
