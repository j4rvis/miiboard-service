package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	delivery "miiboard-service/delivery/http"
	"miiboard-service/repository"
	"miiboard-service/repository/sql"
	"miiboard-service/usecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var (
		apiPort    = os.Getenv("API_PORT")
		dbHost     = os.Getenv("DB_HOST")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbName     = os.Getenv("DB_NAME")
	)
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

	log.Println(apiPort)

	repository := repository.NewDashboardRepository(sql.NewDB(dbHost, dbUser, dbPassword, dbName))
	usecase := usecase.DashboardUseCase(repository)
	dashboardHandler := delivery.NewDashboardHandler(usecase)

	router.Route("/", func(r chi.Router) {
		r.MethodFunc(http.MethodGet, "/dashboard/{id}", dashboardHandler.GetDashboard)
		r.MethodFunc(http.MethodGet, "/health", delivery.Health)
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", apiPort),
		Handler: router,
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
