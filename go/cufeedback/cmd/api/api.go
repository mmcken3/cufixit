package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello world from api")
	router := mux.NewRouter()

	router.HandleFunc("/", GetIndex).Methods("GET")

	log.Fatal(http.ListenAndServe(":8002", router))
}
