package main

import (
	"net"
)

type PingPongService interface {
	Ping(PingRequest) (string, error)
}

type pingPongService struct{}

func (pingPongService) Ping(pingRequest PingRequest) (string, error) {

	remoteAddress := pingRequest.HttpRequest.RemoteAddr

	// Extract the ip from the request Remote Address
	extractedIp, _, splitError := net.SplitHostPort(remoteAddress)

	return extractedIp, splitError
}
