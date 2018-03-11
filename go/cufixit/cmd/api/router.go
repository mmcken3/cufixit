package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mmcken3/cufixit/go/cufixit"
	"github.com/mmcken3/cufixit/go/postgres"
)

// GetIndex returns the index page for the webapp.
func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text")
	data := []byte("Here is a Hello World API Endpoint....")
	_, _ = w.Write(data)
}

// SubmitFixIt parses json from client and saves feedback to the DB.
func SubmitFixIt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userSubmit fixItSubmit
	err := json.NewDecoder(r.Body).Decode(&userSubmit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding submit from json. %v\n", err)
		return
	}

	fixItFeedback := transformToFeedback(userSubmit)
	db, err := postgres.CreateDB()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error starting database. %v\n", err)
		return
	}
	err = db.CreateFeedback(fixItFeedback)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error storing to database. %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	err = db.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error closing database. %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetAllFeedback request all of the feedback in the DB as a json array.
func GetAllFeedback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := postgres.CreateDB()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panicf("Error starting database. %v\n", err)
		return
	}
	feedback, err := db.GetAllFeedback()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panicf("Error getting from database. %v\n", err)
		return
	}
	err = db.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panicf("Error closing database. %v\n", err)
		return
	}

	respondJSON, err := json.Marshal(feedback)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panicf("Error marshalling json. %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respondJSON)
}

type fixItSubmit struct {
	UserName     string `json:"user_name"`
	Type         string `json:"type"`
	BuildingName string `json:"building_name"`
	Description  string `json:"description"`
	PhoneNumber  string `json:"phone_number"`
	ImageURL     string `json:"image_url"`
}

func transformToFeedback(f fixItSubmit) cufixit.Feedback {
	var fb cufixit.Feedback

	fb.UserName = f.UserName
	fb.Description = f.Description
	fb.PhoneNumber = f.PhoneNumber
	fb.ImageURL = f.ImageURL
	fb.Building.Name = f.BuildingName
	fb.Type.Type = f.Type

	return fb
}
