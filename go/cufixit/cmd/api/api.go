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
	router.HandleFunc("/v1/submit", SubmitFixIt).Methods("POST")

	log.Fatal(http.ListenAndServe(":8002", router))
}

/*
	db, err := postgres.CreateDB()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = db.CreateFeedback(cufixit.Feedback{
		ID:       1,
		UserName: "mmcken3",
		Type: cufixit.Type{
			ID:      0,
			Type:    "Housing",
			Contact: "",
		},
		Location:    "Hardin Hall",
		Description: "A table in room 213 of Hardin Hall is broken.",
		PhoneNumber: "843-124-3258",
		ImageURL:    "ins3.com",
		Building: cufixit.Building{
			ID:   1,
			Name: "Hardin Hall",
		},
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	feedback, err := db.GetAllFeedback()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, f := range feedback {
		fmt.Printf("Feedback: %v\n", f)
	}
	db.Close()
*/
