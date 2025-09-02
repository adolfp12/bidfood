package main

import (
	"bidfood/internal/router"
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {

	httpRouter := router.New()

	port := ":8080"
	svr := http.Server{
		Addr:    port,
		Handler: httpRouter,
	}
	go func() {
		log.Printf("Server start to run with port %s", port)
		if err := svr.ListenAndServe(); err != nil {
			log.Printf("Server stopped listening: %v", err)
		}
	}()

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	<-shutdown.Done()

	log.Println("Shutting down server...")

	if err := svr.Shutdown(context.Background()); err != nil {
		log.Printf("Server stopped listening: %v", err)
	}
	log.Println("Shutting down server completed!!")

}
