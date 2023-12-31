package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/api/time", GetTime).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
