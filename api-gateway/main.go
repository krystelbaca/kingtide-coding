package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Clients *GRPCClients
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a Product struct
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Here, you would typically call your gRPC client method to create the product
	// For demonstration purposes, we'll just return the received product in the response
	responseJSON, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Failed to encode product response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

func main() {
	// Initialize gRPC clients
	clients := NewGRPCClients("localhost:50051", "localhost:50052")

	handler := &Handler{Clients: clients}

	r := mux.NewRouter()

	r.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	// r.HandleFunc("/products/{id}", handler.GetProduct).Methods("GET")
	// r.HandleFunc("/orders", handler.PlaceOrder).Methods("POST")
	// r.HandleFunc("/orders", handler.ListOrders).Methods("GET")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
