package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Trick struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Progress string    `json:"progess"`
	State    string    `json:"state"`
	Added_on time.Time `json:"added_on"`
}

type Tricks []Trick

func listAllTricks(w http.ResponseWriter, r *http.Request) {
	tricks := Tricks{
		Trick{Id: 1, Name: "Sit", Progress: "learned", State: "done", Added_on: time.Now()},
		Trick{Id: 2, Name: "Lay down", Progress: "learned", State: "done", Added_on: time.Now()},
	}
	json.NewEncoder(w).Encode(tricks)
}

func insertTrick(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func startPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Start page hit!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", startPage)
	myRouter.HandleFunc("/tricks", listAllTricks).Methods("GET")
	myRouter.HandleFunc("/tricks", insertTrick).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
