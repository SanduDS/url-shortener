package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SanduDS/url-shortener/internal/handler"
	"github.com/SanduDS/url-shortener/internal/service"
)

func main() {
	mux := http.NewServeMux()
	urlService := service.NewURLService()
	urlHandler := handler.NewURLHandler(urlService)

	mux.HandleFunc("/api/v1/shorten", urlHandler.ShortenURL)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "URL Shortener Service is healthy & running 🚀")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server is listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
