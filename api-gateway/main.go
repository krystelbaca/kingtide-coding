package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/products", handleProducts).Methods("GET", "POST")
	r.HandleFunc("/products/{id}", handleProduct).Methods("GET", "PUT", "DELETE")

	r.HandleFunc("/orders", handleOrders).Methods("GET", "POST")
	r.HandleFunc("/orders/{id}", handleOrder).Methods("GET", "PUT", "DELETE")

	http.ListenAndServe(":8080", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API Gateway is up and running"))
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
}

func handleProduct(w http.ResponseWriter, r *http.Request) {
}

func handleOrders(w http.ResponseWriter, r *http.Request) {
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
}
