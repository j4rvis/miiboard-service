package main

import (
	"context"
	"log"
	"math/rand"
	delivery "miiboard-service/delivery/http"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

func main() {
	ctx, done := context.WithCancel(context.Background())
	rand.Seed(time.Now().UnixNano())
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-signalChannel
		log.Printf("signal caught: %v", sig)
		done()
	}()

	router := chi.NewRouter()

	// usecase := usecase.DashboardUseCase(nil)
	dashboardHandler := delivery.NewDashboardHandler(nil)

	router.Route("/", func(r chi.Router) {
		r.MethodFunc(http.MethodGet, "/dashboard/{id}", dashboardHandler.GetDashboard)
		r.MethodFunc(http.MethodGet, "/health", delivery.Health)
	})

	server := http.Server{
		Handler: router,
		Addr:    ":8081",
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("error starting service: %v", err)
			done()
		}
	}()
	log.Println("service started")
	<-ctx.Done()
	log.Println("service stopped")
}
