package main

type PingPongService interface {
	Ping() (string, error)
}

type pingPongService struct{}

func (pingPongService) Ping() (string, error) {
	return "pong", nil
}
