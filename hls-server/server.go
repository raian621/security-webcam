package main

import (
	"net"
	"net/http"
)

type HlsServer struct {
	Server     http.Server
	StopStream chan struct{}
}

func NewHlsServer(addr, port string) *HlsServer {
	mux := http.NewServeMux()
	addHandlers(mux)

	return &HlsServer{
		Server: http.Server{
			Addr:    net.JoinHostPort(addr, port),
			Handler: mux,
		},
		StopStream: make(chan struct{}),
	}
}
