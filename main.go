// Main contains the entire K-Drama CRUD API. It creates the structs, CRUD functions, and the endpoints.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Drama struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Writer   *Writer   `json:"writer"`
	Director *Director `json:"director"`
}

type Writer struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var dramas []Drama

// Returns the entire dramas slice as a json object
func getDramas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dramas)
}

// Deletes a drama struct from the dramas slice by appending the preceding structs and the following structs
// i.e. excluding the desired drama struct from the append statement
func deleteDrama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range dramas {
		if item.ID == params["id"] {
			dramas = append(dramas[:index], dramas[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(dramas)
}

// Gets a particular drama by ID and returns the corresponding json object
func getDrama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range dramas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Creates a new drama struct, appends it to the dramas slice, and returns the new json object
func createDrama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var drama Drama

	// Here we are decoding the json drama object received in the POST request
	_ = json.NewDecoder(r.Body).Decode(&drama)

	// Set new internal ID then append to dramas slice
	drama.ID = strconv.Itoa(rand.Intn(100000))
	dramas = append(dramas, drama)

	json.NewEncoder(w).Encode(drama)
}

// Updates a drama by deleting it from the dramas slice and appending a new version.
func updateDrama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range dramas {
		if item.ID == params["id"] {
			dramas = append(dramas[:index], dramas[index+1:]...)
			var drama Drama
			_ = json.NewDecoder(r.Body).Decode(&drama)
			drama.ID = params["id"]
			dramas = append(dramas, drama)
			json.NewEncoder(w).Encode(drama)
		}
	}
}

func main() {
	// Create new router instance
	r := mux.NewRouter()

	// Create two initial dramas, so that GET methods can be immediately performed
	dramas = append(dramas, Drama{ID: "1", Isbn: "995670", Title: "Start-Up",
		Writer:   &Writer{Firstname: "Hye-ryun", Lastname: "Park"},
		Director: &Director{Firstname: "Choong-hwan", Lastname: "Oh"}})

	dramas = append(dramas, Drama{ID: "2", Isbn: "789456", Title: "Sky Castle",
		Writer:   &Writer{Firstname: "Hyun-mi", Lastname: "Yoo"},
		Director: &Director{Firstname: "Hyun-tak", Lastname: "Jo"}})

	// Define API endpoints & set their functions
	r.HandleFunc("/api/v1/dramas", getDramas).Methods("GET")
	r.HandleFunc("/api/v1/dramas/{id}", getDrama).Methods("GET")
	r.HandleFunc("/api/v1/dramas", createDrama).Methods("POST")
	r.HandleFunc("/api/v1/dramas/{id}", updateDrama).Methods("PUT")
	r.HandleFunc("/api/v1/dramas/{id}", deleteDrama).Methods("DELETE")

	// Start server on port 8000 and catch any fatal errors
	fmt.Printf("Starting server at port 8000...\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
