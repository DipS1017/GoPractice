package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snipperCreate)

	log.Print("Server is running on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
