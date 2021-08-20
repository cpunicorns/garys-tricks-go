package garys_tricks

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", startPage)
	myRouter.HandleFunc("/tricks", listAllTricks).Methods("GET")
	myRouter.HandleFunc("/tricks", insertTrick).Methods("POST")
	myRouter.HandleFunc("/tricks/{id}", listOneTrick).Methods("GET")
	//myRouter.HandleFunc("/tricks/{id}", updateOneTrick).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
