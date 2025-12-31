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

	mux.HandleFunc("POST /trip/preview", enableCORS(handleTripPreview))
	// ws
	mux.HandleFunc("/ws/drivers", handleDriversWebSocket)
	mux.HandleFunc("/ws/riders", handleRidersWebSocket)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	servereErrors := make(chan error, 1)

	go func() {
		log.Panicf("Server listening on: %s", httpAddr)
		servereErrors <- server.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-stop:
		log.Panicf("Server is shutting down due to %v signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("falied to shutdonw gracefully: %v", err)
			server.Close()
		}
	case err := <-servereErrors:
		log.Printf("Error starting the server: %v", err)
	}
}
