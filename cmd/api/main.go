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
		"POST /api/outlets/post",
		handler.CreateOutlet(db),
	)

	mux.HandleFunc(
		"GET /api/outlets/get",
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
		"POST /api/payments/post",
		handler.CreatePayment(db),
	)

	mux.HandleFunc(
		"GET /api/payments/get",
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc(
		"POST /api/products/post",
		handler.CreateProduct(db),
	)

	mux.HandleFunc(
		"GET /api/products/get",
		handler.ReadProducts(db),
	)

	mux.HandleFunc(
		"GET /api/products/outlet/{id}",
		handler.ReadProductsByOutlet(db),
	)

	mux.HandleFunc(
		"GET /api/products/available/{available}",
		handler.ReadProductsByAvailable(db),
	)

	mux.HandleFunc(
		"PATCH /api/products/{id}/price",
		handler.UpdateProductPrice(db),
	)

	mux.HandleFunc(
		"PATCH /api/products/{id}/available",
		handler.UpdateProductAvailable(db),
	)

	mux.HandleFunc(
		"POST /api/sale-items/post",
		handler.CreateSaleItem(db),
	)

	mux.HandleFunc(
		"GET /api/sale-items/get",
		handler.ReadSaleItems(db),
	)

	mux.HandleFunc(
		"GET /api/sale-items/sale/{id}",
		handler.ReadSaleItemsBySale(db),
	)

	mux.HandleFunc(
		"GET /api/sale-items/product/{id}",
		handler.ReadSaleItemsByProduct(db),
	)

	mux.HandleFunc(
		"POST /api/sales/post",
		handler.CreateSale(db),
	)

	mux.HandleFunc(
		"GET /api/sales/get",
		handler.ReadSales(db),
	)

	mux.HandleFunc(
		"GET /api/sales/user/{id}",
		handler.ReadSalesByUser(db),
	)

	mux.HandleFunc(
		"GET /api/sales/outlet/{id}",
		handler.ReadSalesByOutlet(db),
	)

	mux.HandleFunc(
		"GET /api/sales/date/{date}",
		handler.ReadSalesByDate(db),
	)

	mux.HandleFunc(
		"PATCH /api/sales/status/{id}",
		handler.UpdateSaleStatus(db),
	)

	mux.HandleFunc(
		"POST /api/stock",
		handler.CreateStock(db),
	)

	mux.HandleFunc(
		"GET /api/stock",
		handler.ReadStock(db),
	)

	mux.HandleFunc(
		"GET /api/stock/outlet/{id}",
		handler.ReadOutletStock(db),
	)

	mux.HandleFunc(
		"PATCH /api/stock/{id}/quantity",
		handler.UpdateStockQty(db),
	)

	mux.HandleFunc(
		"PATCH /api/stock/{id}/price",
		handler.UpdateStockPrice(db),
	)

	mux.HandleFunc("POST /api/workers/post", handler.CreateWorker(db))

	mux.HandleFunc("GET /api/workers/get", handler.GetWorkers(db))

	mux.HandleFunc(
		"GET /api/workers/outlet/{id}",
		handler.GetWorkersByOutlet(db),
	)

	mux.HandleFunc(
		"POST /api/workers/signin/{id}",
		handler.SignInWorker(db),
	)

	mux.HandleFunc(
		"PATCH /api/workers/shift/{id}",
		handler.UpdateShiftTime(db),
	)

	mux.HandleFunc(
		"PATCH /api/workers/location/{id}",
		handler.UpdateSignInLocation(db),
	)

	mux.HandleFunc(
		"PATCH /api/workers/outlet/{id}",
		handler.UpdateOutletID(db),
	)

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
