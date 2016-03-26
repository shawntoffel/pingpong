package main

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func MakePingEndpoint(service pingPongService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		pingRequest := request.(PingRequest)

		pong, err := service.Ping(pingRequest.RemoteAddress)

		if err != nil {
			return ErrorResponse{"Could not determine client IP address", err.Error()}, nil
		}

		return PingResponse{pong}, nil
	}
}
