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
	var userSubmit fixItSubmit
	err := json.NewDecoder(r.Body).Decode(&userSubmit)
	if err != nil {
		// Send error json back
		log.Panicf("Error decoding submit from json. %v\n", err)
	}

	fixItFeedback := transformToFeedback(userSubmit)
	db, err := postgres.CreateDB()
	if err != nil {
		// Send error json back
		log.Panicf("Error starting database. %v\n", err)
	}
	err = db.CreateFeedback(fixItFeedback)
	if err != nil {
		// Send error json back
		log.Panicf("Error storing to database. %v\n", err)
	}
	err = db.Close()
	if err != nil {
		// Send error json back
		log.Panicf("Error closing database. %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

type fixItSubmit struct {
	UserName     string `json:"user_name"`
	Type         string `json:"type"`
	BuildingName string `json:"building_name"`
	Location     string `json:"location"`
	Description  string `json:"description"`
	PhoneNumber  string `json:"phone_number"`
	ImageURL     string `json:"image_url"`
}

func transformToFeedback(f fixItSubmit) cufixit.Feedback {
	var fb cufixit.Feedback

	fb.UserName = f.UserName
	fb.Location = f.Location
	fb.Description = f.Description
	fb.PhoneNumber = f.PhoneNumber
	fb.ImageURL = f.ImageURL
	fb.Building.Name = f.BuildingName
	fb.Type.Type = f.Type

	return fb
}
