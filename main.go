package main

import (
	"goth-starter/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("GET /", handlers.CountryList)
	mux.HandleFunc("GET /country/{name}", handlers.CountryDetail)
	mux.HandleFunc("GET /search", handlers.SearchCountries)

	log.Println("Starting server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
