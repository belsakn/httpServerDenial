package server

import (
	"log"
	"net/http"
	"time"
)

type Client struct {
	startTime    time.Time
	requestCount int
}

type Clients map[string]*Client

var clientMap Clients

func Server() {
	clientMap = make(Clients)
	log.Println("HTTP Denial-of-Service protection system listening on port 8080")

	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("------")
	r.ParseForm()
	clientId := r.Form.Get("clientId")

	if clientId == "" || clientId == " " {
		log.Printf("Parameter clientId was not provided.")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Println("clientId: ", clientId)

	currentTime := time.Now()
	if _, ok := clientMap[clientId]; ok {
		clientMap[clientId].requestCount++

		elapsedTime := currentTime.Sub(clientMap[clientId].startTime)
		elapsedTimeInt64 := int64(elapsedTime.Seconds() * 1000)

		log.Println("elapsedTime from last request: ", elapsedTimeInt64, "ms")

		if elapsedTimeInt64 >= 5000 {
			clientMap[clientId].startTime = currentTime
			clientMap[clientId].requestCount = 1
			log.Println("Time frime has passed, reset requestCount.")
			w.WriteHeader(http.StatusOK)
			return
		}

		if clientMap[clientId].requestCount <= 5 {
			w.WriteHeader(http.StatusOK)
			return
		}

		if clientMap[clientId].requestCount >= 5 && elapsedTimeInt64 <= 5000 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

	}

	//clients first request
	clientMap[clientId] = &Client{startTime: currentTime, requestCount: 1}
	log.Println("First request for client.")
	w.WriteHeader(http.StatusOK)
}
