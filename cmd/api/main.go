package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mydb "github.com/NigzMaf1a/atlas-hortus-vitae/internal/db"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/handler"
	"github.com/NigzMaf1a/atlas-hortus-vitae/internal/middleware"
)

func main() {
	// Database connection
	db, err := mydb.ConnectDB()

	if err != nil {
		log.Fatalf("opening database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("connecting to database: %v", err)
	}

	log.Println("Database connected.")

	mux := http.NewServeMux()

	//routes
	mux.HandleFunc("POST /api/auth/login", handler.Login(db))

	mux.HandleFunc(
		"POST /api/outlets",
		handler.CreateOutlet(db),
	)

	mux.HandleFunc(
		"GET /api/outlets",
		handler.ReadOutlets(db),
	)

	mux.HandleFunc(
		"GET /api/outlets/{id}",
		handler.ReadOutlet(db),
	)

	mux.HandleFunc(
		"PATCH /api/outlets/{id}/networth",
		handler.UpdateNetworth(db),
	)

	mux.HandleFunc(
		"PATCH /api/outlets/{id}/status",
		handler.UpdateOutletStatus(db),
	)

	mux.HandleFunc(
		"POST /api/payments",
		handler.CreatePayment(db),
	)

	mux.HandleFunc(
		"GET /api/payments",
		handler.ReadPayments(db),
	)

	mux.HandleFunc(
		"GET /api/payments/user/{id}",
		handler.ReadPaymentsByUser(db),
	)

	mux.HandleFunc(
		"PATCH /api/payments/{id}/status",
		handler.UpdatePaymentStatus(db),
	)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Server is running"))
	})

	// Server configuration
	server := &http.Server{
		Addr:              ":" + getPort(),
		Handler:           middleware.CORS(mux),
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Start server
	go func() {
		log.Printf("Server listening on http://localhost%s", server.Addr)

		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully.")
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	return port
}
