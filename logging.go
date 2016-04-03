package main

import (
	"github.com/go-kit/kit/log"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	PingPongService
}

func (mw LoggingMiddleware) HandleLoggingError(err error, method string) error {
	if err != nil {
		defer func(begin time.Time) {
			mw.Logger.Log(
				"loggingError", err,
				"method", method,
				"ts", begin.UTC(),
			)
		}(time.Now())
	}

	return nil
}

func (mw LoggingMiddleware) Ping(pingRequest PingRequest) (output string, err error) {
	defer func(begin time.Time) {
		err = mw.Logger.Log(
			"method", "Ping",
			"requestAddress", pingRequest.HttpRequest.RemoteAddr,
			"output", output,
			"error", err,
			"took", time.Since(begin),
			"ts", begin.UTC(),
		)

		err = mw.HandleLoggingError(err, "Ping")

	}(time.Now())

	output, err = mw.PingPongService.Ping(pingRequest)

	return
}
