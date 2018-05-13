package server

import (
	"fmt"
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

	r.ParseForm()
	clientId := r.Form.Get("clientId")

	if clientId == "" || clientId == " " {
		log.Printf("Parameter clientId was not provided. RETURN 403")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(responseCodeForClient(clientId))
}

func responseCodeForClient(clientId string) int {
	logString := fmt.Sprintf("ClientId: %s > ", clientId)

	currentTime := time.Now()

	//Check if client exists else add client
	if _, ok := clientMap[clientId]; ok {
		clientMap[clientId].requestCount++

		elapsedTime := currentTime.Sub(clientMap[clientId].startTime)
		elapsedTimeInt64 := int64(elapsedTime.Seconds() * 1000)

		logString = logString + fmt.Sprintf("elapsedTime from last request: %d ms. Client request counter: %d. ", elapsedTimeInt64, clientMap[clientId].requestCount)

		//elapsedTime is greater the 5000 than reset counter and start new time frame window from now
		if elapsedTimeInt64 >= 5000 {
			clientMap[clientId].startTime = currentTime
			clientMap[clientId].requestCount = 1
			log.Printf(logString + "Time frime has passed, reset requestCount. RETURN 200\n")
			return http.StatusOK

		}

		//Client made less the 5 request than OK
		if clientMap[clientId].requestCount <= 5 {
			log.Printf(logString + "RETURN 200\n")
			return http.StatusOK
		}

		//Client has reached the threshold of 5 req in 5s so 503 is returned
		if clientMap[clientId].requestCount >= 5 && elapsedTimeInt64 <= 5000 {
			log.Printf(logString + "Too much request for clients time frame. RETURN 503\n")
			return http.StatusServiceUnavailable
		}

	}

	//Clients first request
	clientMap[clientId] = &Client{startTime: currentTime, requestCount: 1}
	log.Printf(logString + "First request for client. RETURN 200\n")
	return http.StatusOK
}
