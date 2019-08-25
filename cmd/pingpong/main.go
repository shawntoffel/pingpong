package main

import (
	"net/http"

	"github.com/shawntoffel/pingpong"
)

func main() {
	http.HandleFunc("/", pingpong.OriginatingIP)
	http.ListenAndServe(":8080", nil)
}
