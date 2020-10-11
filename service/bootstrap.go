package service

import (
	"context"
	"go-book-inventory/repository"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Bootstrap() {
	ctx, done := context.WithCancel(context.Background())
	rand.Seed(time.Now().UnixNano())
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-signalChannel
		log.Printf("signal caught: %v", sig)
		done()
	}()

	c:=NewController(*repository.NewBookRepository())

	rte := NewRouter()
	rte.AddRoute(http.MethodGet, "/health", c.healthHandler)
	rte.AddRoute(http.MethodPost,"/books",c.booksHandler)

	server := http.Server{
		Addr:    ":8000",
		Handler: rte,
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
