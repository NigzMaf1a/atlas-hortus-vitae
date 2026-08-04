package router

import (
	"net/http"
	// "github.com/NigzMaf1a/atlas-hortus-vitae/internal/middleware"
)

func SetupRouter() http.Handler {

	mux := http.NewServeMux()

	// // Public routes
	// mux.HandleFunc("/health", Health)

	// // Protected routes
	// mux.Handle(
	// 	"/api/core",
	// 	middleware.Authenticate(
	// 		http.HandlerFunc(CoreHandler),
	// 	),
	// )

	return mux
}
