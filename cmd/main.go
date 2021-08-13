package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Trick struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Progress string    `json:"progress"`
	State    string    `json:"state"`
	Added_on time.Time `json:"added_on"`
}

type Tricks []Trick

var AllTricks = []Trick{
	{Id: "1", Name: "Sit", Progress: "learned", State: "done", Added_on: time.Now()},
	{Id: "2", Name: "Lay down", Progress: "learned", State: "done", Added_on: time.Now()},
}

func listAllTricks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(AllTricks)
}

func listOneTrick(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, trick := range AllTricks {
		if trick.Id == key {
			json.NewEncoder(w).Encode(trick)
		}
	}
}

func insertTrick(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var trick Trick
	json.Unmarshal(reqBody, &trick)
	AllTricks = append(AllTricks, trick)

	json.NewEncoder(w).Encode(trick)
}

func startPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Start page hit!")
}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", startPage)
	myRouter.HandleFunc("/tricks", listAllTricks).Methods("GET")
	myRouter.HandleFunc("/tricks", insertTrick).Methods("POST")
	myRouter.HandleFunc("/tricks/{id}", listOneTrick).Methods("GET")
	//myRouter.HandleFunc("/tricks/{id}", updateOneTrick).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	HandleRequests()
}
