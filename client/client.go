package client

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Client() {

	for {
		makeGetRequest("http://localhost:8080/?clientId=2")
		r := rand.Intn(1222)
		time.Sleep(time.Duration(r) * time.Millisecond)
	}

}

func makeGetRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	log.Printf("Get %s response code is %d", url, resp.StatusCode)
}
