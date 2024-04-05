package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Year     int       `json:"year"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie
var nextID = 1

// getMovies function
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// getMovie function
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createMovie function
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	if len(movies) == 0 {
		movie.ID = strconv.Itoa(nextID)
		movies = append(movies, movie)
		json.NewEncoder(w).Encode(movie)
		return
	}
	lastID, _ := strconv.Atoi(movies[len(movies)-1].ID)
	nextID := lastID + 1
	movie.ID = strconv.Itoa(nextID)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// updateMovie function
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Loop over the movies
	for index, item := range movies {
		if item.ID == params["id"] {
			// Update the existing movie's fields
			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = params["id"] // Keep the same ID
			movies[index] = updatedMovie
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
}

// deleteMovie function
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	// Test
	// movies = append(movies, Movie{ID: "1", Title: "Parasite", Year: 2019, Director: &Director{FirstName: "Joon-ho", LastName: "Bong"}})
	// movies = append(movies, Movie{ID: "2", Title: "Bacurau", Year: 2019, Director: &Director{FirstName: "Kleber", LastName: "Mendonça Filho"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
