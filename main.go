package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	// go func() {
	// 	err := server.ListenAndServe()
	// 	if err != nil {
	// 		log.Printf("error starting service: %v", err)
	// 		done()
	// 	}
	// }()
	log.Println("service started")
	<-ctx.Done()
	log.Println("service stopped")
}
