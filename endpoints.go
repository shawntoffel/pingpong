package main

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func MakePingEndpoint(service PingPongService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		pingRequest := request.(PingRequest)

		pong, err := service.Ping(pingRequest)

		if err != nil {
			return ErrorResponse{"Could not determine IP address", err.Error()}, nil
		}

		return PingResponse{pong}, nil
	}
}
