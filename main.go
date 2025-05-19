package main

import (
	"goth-starter/handlers"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	// Initialize router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Routes
	r.Get("/", handlers.MoviesList)
	r.Get("/movie/{id}", handlers.MovieDetail)

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
