package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {
	log.Println("Starting API Gateway")

	inmemRepo := repository.NewInmemRepository()

	svc := service.NewService(inmemRepo)

	httpHandler := h.NewHttpHandler(*svc)
	mux := http.NewServeMux()

	mux.HandleFunc("Post /preview", httpHandler.HandleTripPreview)

	server := http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	servereError := make(chan error, 1)
	servereError <- server.ListenAndServe()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Println("falied to shutdonw gracefully.")
			server.Close()
		}
	case <-servereError:
	}
}
