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

	ipAddress, _, err := net.SplitHostPort(remoteAddress)

	return ipAddress, err
}
