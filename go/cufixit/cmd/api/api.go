package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mmcken3/cufixit/go/cufixit"

	"github.com/gorilla/mux"
	"github.com/mmcken3/cufixit/go/postgres"
)

func main() {
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
	router := mux.NewRouter()

	router.HandleFunc("/", GetIndex).Methods("GET")

	log.Fatal(http.ListenAndServe(":8002", router))
}
