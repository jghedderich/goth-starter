package handlers

import (
	"encoding/json"
	"fmt"
	"goth-starter/components"
	"goth-starter/models"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}

func MoviesList(w http.ResponseWriter, r *http.Request) {
	baseURL := getRequiredEnv("TMDB_BASE_URL")
	apiKey := getRequiredEnv("TMDB_API_KEY")
	url := fmt.Sprintf("%s/movie/popular?api_key=%s", baseURL, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	var movieResponse models.MovieListResponse
	err = json.NewDecoder(resp.Body).Decode(&movieResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Construct full poster URL (TMDB returns relative paths)
	for i := range movieResponse.Results {
		movieResponse.Results[i].PosterPath = "https://image.tmdb.org/t/p/w500" + movieResponse.Results[i].PosterPath
	}

	component := components.MoviesPage(movieResponse)
	component.Render(r.Context(), w)
}

func MovieDetail(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "id")
	baseURL := getRequiredEnv("TMDB_BASE_URL")
	apiKey := getRequiredEnv("TMDB_API_KEY")

	url := fmt.Sprintf("%s/movie/%s?api_key=%s", baseURL, movieID, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var movieResponse models.MovieDetailResponse
	err = json.NewDecoder(resp.Body).Decode(&movieResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	movieResponse.PosterPath = "https://image.tmdb.org/t/p/w500" + movieResponse.PosterPath

	component := components.MovieDetail(movieResponse)
	component.Render(r.Context(), w)
}

func SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	baseURL := getRequiredEnv("TMDB_BASE_URL")
	apiKey := getRequiredEnv("TMDB_API_KEY")
	var url string
	if query == "" {
		url = fmt.Sprintf("%s/discover/movie?api_key=%s", baseURL, apiKey)
	} else {
		url = fmt.Sprintf("%s/search/movie?api_key=%s&query=%s", baseURL, apiKey, query)
	}
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	var movieResponse models.MovieListResponse
	err = json.NewDecoder(resp.Body).Decode(&movieResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Construct full poster URL (TMDB returns relative paths)
	for i := range movieResponse.Results {
		movieResponse.Results[i].PosterPath = "https://image.tmdb.org/t/p/w500" + movieResponse.Results[i].PosterPath
	}

	component := components.MovieList(movieResponse)
	component.Render(r.Context(), w)
}
