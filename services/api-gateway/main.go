package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", handleTripPreview)

	server := http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	servereError := make(chan error, 1)
	go func() {
		log.Panicf("Server is listening on: %s", httpAddr)
		servereError <- server.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("falied to shutdonw gracefully: %v", err)
			server.Close()
		}
	case err := <-servereError:
		log.Printf("Error starting the server: %v", err)
	}
}
