package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	// Your code here
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Your code here
	w.Header().Set("Content-Type", "application/json")
	// Get the movie ID from the URL parameters
	params := mux.Vars(r)
	id := params["id"]
	fmt.Printf("id: %s", id)

	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// Your code here
	w.Header().Set("Content-Type", "application/json")
	// Get the movie ID from the URL parameters
	params := mux.Vars(r)
	id := params["id"]
	fmt.Printf("id: %s", id)

	for _, item := range movies {
		if item.ID == id {

			json.NewEncoder(w).Encode(item)

			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// Your code here
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Your code here
	fmt.Printf("update movie called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var movie Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		fmt.Println("error")
		http.Error(w, err.Error(), http.StatusBadRequest)

	}
	for index, item := range movies {
		if item.ID == id {
			fmt.Println(movies[index])
			movies[index].Director = movie.Director
			movies[index].Title = movie.Title
			movies[index].ID = movie.ID
			movies[index].Isbn = movie.Isbn
			// movies[index] = movie
		}
	}
	json.NewEncoder(w).Encode(movie)

}

func main() {
	// Your code here
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "The Godfather", Director: &Director{Id: "1", Firstname: "Francis", Lastname: "Coppola"}})
	movies = append(movies, Movie{ID: "2", Isbn: "454555", Title: "The Dark Knight", Director: &Director{Id: "2", Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "3", Isbn: "123456", Title: "Inception", Director: &Director{Id: "3", Firstname: "Christopher", Lastname: "Nolan"}})
	fmt.Printf("random int %T", strconv.Itoa(rand.Intn(10000000)))
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server started on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
