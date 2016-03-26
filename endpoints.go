package main

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func MakePingEndpoint(service pingPongService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		pong, err := service.Ping()

		if err != nil {
			return ErrorResponse{err.Error(), ""}, nil
		}

		return PingResponse{pong}, nil
	}
}
