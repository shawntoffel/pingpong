package main

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Pong string `json:"pong"`
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Detail string `json:"detail"`
}

func DecodePingRequest(r *http.Request) (request interface{}, err error) {
	return
}

func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
