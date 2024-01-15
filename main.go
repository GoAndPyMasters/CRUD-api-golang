package main

// In this project we will be building a  basic API using Go and gorilla/mux that allows users to create, read, update, and delete movies.
// here I am importing all this the nessaery libraries
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// this is the struct that we will be using to create, read, update, and delete movies
type Movie struct {
	// this is the ID of the movie and i am setting this way so i can convert it to JSON
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// this is the struct for Director
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// this is the variable that we will be using to store the movies
var movies []Movie

// getMovies is a function that handles the HTTP request for getting movies.
//
// It takes in two parameters: w http.ResponseWriter and r *http.Request.
// The function does not return any value.
func getMovies(w http.ResponseWriter, r *http.Request) {
	// set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// deleteMovie deletes a movie from the movies list.
//
// It takes in a http.ResponseWriter and a *http.Request as parameters.
// It does not return anything.
func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// params contains the ID of the movie to be deleted and parsed from the URL or headers
	params := mux.Vars(r)
	// loop through the list of movies
	for index, item := range movies {
		if item.ID == params["id"] {
			// overwrite the movie at the index with the next movie in the list
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// encode the all movies into JSON format and send it back to the client
	json.NewEncoder(w).Encode(movies)
}

// getMovie retrieves a movie by ID and encodes it into JSON format.
//
// The function takes in two parameters:
//   - w: the http.ResponseWriter used to write the HTTP response.
//   - r: the *http.Request representing the HTTP request.
//
// The function does not return any values.
func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// params contains the ID of the movie to be retrieved and parsed from the URL or headers
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			// encode the movie into JSON format and send it back to the client
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// encode the all movies into JSON format and send it back to the client
	json.NewEncoder(w).Encode(&Movie{})
}

// createMovie handles the creation of a new movie.
//
// It takes in two parameters:
// - w: an http.ResponseWriter object used to write the HTTP response.
// - r: an *http.Request object representing the incoming HTTP request.
//
// This function does not return anything.
func createMovie(w http.ResponseWriter, r *http.Request) {
	// set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// create a new movie object
	var movie Movie
	// decode the request body into the movie object
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// generate a unique ID for the movie
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	// append the new movie to the list of movies
	movies = append(movies, movie)
	// encode the new movie into JSON format and send it back to the client
	json.NewEncoder(w).Encode(movie)
}

// updateMovie updates a movie in the movies slice based on the ID provided in the request parameters.
//
// It takes in two parameters:
// - w: an http.ResponseWriter used to write the response back to the client.
// - r: a pointer to an http.Request object representing the incoming request.
//
// It does not return any value.
func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	fmt.Print("Movie is upaded")
	json.NewEncoder(w).Encode(movies)
}

// main is the entry point function for the program.
//
// It initializes the movies slice with two movie objects.
// It creates a new router using the mux package.
// It registers the HTTP handlers for various movie routes.
// It starts the HTTP server on port 8000.
func main() {
	// create a slice of movies
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
	// create a new router
	r := mux.NewRouter()
	// register the routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// register the route for a single movie
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// register the route for creating a new movie
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// register the route for updating a movie
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// register the route for deleting a movie
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	// start the server
	fmt.Printf("String server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
