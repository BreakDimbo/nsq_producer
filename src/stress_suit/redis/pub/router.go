package main

import (
	"net/http"
	"stress_suit/redis/pub/action"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/redis/pub/topics", action.Mad).Methods("POST")
	router.HandleFunc("/redis/pub/stop", action.Stop).Methods("POST")
	return router
}
