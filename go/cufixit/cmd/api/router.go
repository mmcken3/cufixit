package main

import (
	"net/http"
)

// GetIndex returns the index page for the webapp.
func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text")
	data := []byte("Here is a Hello World API Endpoint....")
	_, _ = w.Write(data)
}
