package main

import (
	"log"
	"net/http"

	"github.com/THETITAN220/auto-scale-URL/internal/handler"
	"github.com/THETITAN220/auto-scale-URL/internal/service"
	"github.com/THETITAN220/auto-scale-URL/internal/storage"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	store := storage.NewMemoryStore()
	svc := service.NewURLService(store)
	h := handler.NewHandler(svc)

	http.HandleFunc("/shorten", h.Shorten)
	http.HandleFunc("/r/", h.Redirect)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}