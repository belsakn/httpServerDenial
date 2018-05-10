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
	r.ParseForm()
	clientId := r.Form.Get("clientId")

	if clientId == "" || clientId == " " {
		log.Printf("Parameter clientId was not provided.")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Printf("Phone number to be parsed %s", clientId)

	w.WriteHeader(http.StatusOK)
}
