package main

import (
	"log"
	"net/http"

	"github.com/AndreiMartynenko/crypto-payment-gateway/internal/middleware"
	"github.com/AndreiMartynenko/crypto-payment-gateway/internal/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// Apply middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)

	// Define routes
	r.HandleFunc("/initiate-payment", routes.HandleInitiatePayment).Methods("POST")
	r.HandleFunc("/verify-transaction", routes.HandleVerifyTransaction).Methods("GET")

	// Start the server
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
	handler := c.Handler(r)

	log.Println("API Gateway is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
