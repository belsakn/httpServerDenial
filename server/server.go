package server

import (
	"log"
	"net/http"
)

func Server() {
	log.Println("HTTP Denial-of-Service protection system listening on port 8080")

	http.HandleFunc("/", handlerequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
