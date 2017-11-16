package main

import (
	"net/http"
	"nsq_stress/nsq_producer/action"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/producer/mad", action.Mad).Methods("GET")
	router.HandleFunc("/producer/stop", action.Stop).Methods("GET")
	return router
}
