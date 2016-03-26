package main

import (
	"github.com/go-kit/kit/log"
	"net/http"
	"os"

	"golang.org/x/net/context"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func main() {
	backgroundContext := context.Background()
	logger := log.NewJSONLogger(os.Stderr)

	var service PingPongService
	service = pingPongService{}
	service = LoggingMiddleware{logger, service}

	pongHandler := httpTransport.NewServer(
		backgroundContext,
		MakePingEndpoint(service),
		DecodePingRequest,
		EncodeResponse,
	)

	http.Handle("/ping", pongHandler)
	logger.Log("msg", "HTTP", "addr", ":10002")
	logger.Log("err", http.ListenAndServe(":10003", nil))
}
