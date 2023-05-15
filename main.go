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
	ID       string    `json :"id"`
	Isbn     string    `json : "isbn"`
	Title    string    `jason : "title"`
	Director *Director `json  : "director"`
}

type Director struct {
	Firstname string `json:"fristname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	movies = append(movies, Movie{ID: "238", Isbn: "438565", Title: "Avengers infinitywar", Director: &Director{Firstname: "hary ", Lastname: "potter"}})
	movies = append(movies, Movie{ID: "241", Isbn: "438227", Title: "RRR", Director: &Director{Firstname: "SS", Lastname: "rajamoli"}})
	movies = append(movies, Movie{ID: "630", Isbn: "438226", Title: "lengend of prince ram", Director: &Director{Firstname: "trgerg", Lastname: "hfthvd"}})
	movies = append(movies, Movie{ID: "431", Isbn: "438234", Title: "The king"})

	r := mux.NewRouter()

	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movies", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")

	fmt.Println("starting server at port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))

}

func getmovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content_type", "application/json")
	json.NewEncoder(w).Encode(movies)

}
func deletemovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content_type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)
}

func getmovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content_type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return

		}

	}

}

func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content_type", "application/json")

	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(999))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(r)

}

func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content_type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(r)
			return
		}

	}

}
