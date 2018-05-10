package client

import (
	"log"
	"net/http"
)

func Client() {

	urls := []string{"http://localhost:8080/", "http://localhost:8080/?clientId=2", "http://localhost:8080/fdgd"}

	for _, url := range urls {
		makeGetRequest(url)
	}

}

func makeGetRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	log.Printf("Get %s response code is %d", url, resp.StatusCode)
}
