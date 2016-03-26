package main

import (
	"net"
)

type PingPongService interface {
	Ping(string) (string, error)
}

type pingPongService struct{}

func (pingPongService) Ping(remoteAddress string) (string, error) {
	// Extract the ip from the request Remote Address
	extractedIp, _, splitError := net.SplitHostPort(remoteAddress)

	return extractedIp, splitError
}
