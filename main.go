package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type HTTPHandler struct {
	ModernService ModernService
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func (handler *HTTPHandler) SubmitOrderHandler(w http.ResponseWriter, req *http.Request) {
	var order *Order
	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		w.Write([]byte("Payload malformed."))
	}

	orderCreated, err := handler.ModernService.CreateOrder(order)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		// I know this is not ideal, but I'll do it
		w.Write([]byte(err.Error()))
	}
	if orderCreated {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		w.Write([]byte("Order created successfuly!"))
	}
}

func main() {
	var handler HTTPHandler

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/order", handler.SubmitOrderHandler)
	slog.Info("Service started in port 8000")
	http.ListenAndServe(":8000", nil)
}
