package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	"github.com/shawntoffel/pingpong/controllers"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a PongController instance
	pongController := controllers.NewPongController()

	// Route
	r.GET("/ping", pongController.GetPong)

	// Serve
	log.Fatal(http.ListenAndServe(":10002", r))
}
