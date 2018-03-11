package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetIndex).Methods("GET")
	router.HandleFunc("/v1/getall", GetAllFeedback).Methods("GET")
	router.HandleFunc("/v1/get/type/{type}", GetFeedbackofType).Methods("GET")
	router.HandleFunc("/v1/get/building/{building}", GetFeedbackofBuilding).Methods("GET")
	router.HandleFunc("/v1/get/user/{user}", GetFeedbackofUser).Methods("GET")
	router.HandleFunc("/v1/submit", SubmitFixIt).Methods("POST")

	log.Fatal(http.ListenAndServe(":8002", router))
}
