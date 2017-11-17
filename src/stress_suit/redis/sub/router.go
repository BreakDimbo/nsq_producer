package main

import (
	"net/http"
	"stress_suit/redis/sub/action"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/redis/sub/topics", action.Mad).Methods("POST")
	router.HandleFunc("/redis/sub/stop", action.Stop).Methods("POST")
	router.HandleFunc("/redis/sub/set", action.Set).Methods("POST")
	return router
}
