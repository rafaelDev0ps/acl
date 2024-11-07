package main

import (
	"log/slog"
	"net/http"
)

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/health", healthHandler)
	slog.Info("Service started in port 8000")
	http.ListenAndServe(":8000", nil)
}
