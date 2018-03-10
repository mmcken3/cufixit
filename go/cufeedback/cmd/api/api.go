package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mmcken3/cufeedback/go/cufeedback"

	"github.com/gorilla/mux"
	"github.com/mmcken3/cufeedback/go/postgres"
)

func main() {
	db, err := postgres.CreateDB()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = db.CreateFeedback(cufeedback.Feedback{
		ID:          1,
		UserName:    "mmcken3",
		Type:        "Broken Table",
		Location:    "Hardin Hall",
		Description: "A table in room 213 of Hardin Hall is broken.",
		Email:       "facilities.email",
		Building: cufeedback.Building{
			ID:   1,
			Name: "Hardin Hall",
		},
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	db.Close()
	router := mux.NewRouter()

	router.HandleFunc("/", GetIndex).Methods("GET")

	log.Fatal(http.ListenAndServe(":8002", router))
}
