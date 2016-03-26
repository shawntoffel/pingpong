package main

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func main() {
	backgroundContext := context.Background()
	service := pingPongService{}

	pongHandler := httpTransport.NewServer(
		backgroundContext,
		MakePingEndpoint(service),
		DecodePingRequest,
		EncodeResponse,
	)

	http.Handle("/ping", pongHandler)
	log.Fatal(http.ListenAndServe(":10002", nil))
}
