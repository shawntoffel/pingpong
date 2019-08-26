package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shawntoffel/pingpong"
)

// Version of pingpong
var Version = ""

var (
	flagVersion = false
	flagPort    = "8080"
)

func parseCli() {
	flag.BoolVar(&flagVersion, "v", false, "version")
	flag.StringVar(&flagPort, "p", flagPort, "port")

	port := os.Getenv("PORT")
	if port != "" {
		flagPort = port
	}

	flag.Parse()
}

func main() {
	parseCli()

	if flagVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	log.Printf("%s starting", Version)

	server := http.Server{
		Addr: ":" + flagPort,
	}

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

		<-signals

		log.Printf("stopping")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	http.HandleFunc("/", pingpong.OriginatingIPHandler)
	log.Print(server.ListenAndServe())
}
