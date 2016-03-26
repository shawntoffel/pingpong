package main

import (
	"github.com/go-kit/kit/log"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	PingPongService
}

func (mw LoggingMiddleware) Ping(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"Method", "Ping",
			"RequestAddress", s,
			"Output", output,
			"Error", err,
			"Took", time.Since(begin),
		)

	}(time.Now())

	output, err = mw.PingPongService.Ping(s)

	return
}
