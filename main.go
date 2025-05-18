package main

import (
	"goth-starter/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Routes
	r.Get("/", handlers.CountryList)
	r.Get("/country/{name}", handlers.CountryDetail)
	r.Get("/search", handlers.SearchCountries)

	log.Println("Starting server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
