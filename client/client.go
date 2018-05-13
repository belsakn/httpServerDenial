package client

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup
var stopWorking bool = false

func Client(numberOfClinets int) {

	setupCloseHandler()

	for i := 1; i <= numberOfClinets; i++ {
		wg.Add(1)
		go startClient(fmt.Sprintf("http://localhost:8080/?clientId=%d", i), i)
	}

	wg.Wait()
	log.Printf("My work have is done...")
}

func startClient(url string, id int) {
	for {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		log.Printf("Get %s response code is %d", url, resp.StatusCode)

		if stopWorking {
			break
		}

		r := rand.Intn(1222)
		time.Sleep(time.Duration(r) * time.Millisecond)
	}
	fmt.Println("# Client with id: ", id, " stoped making requests.")
	wg.Done()
}

func setupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("# Wait for all Clients to finish their requests...")
		stopWorking = true
	}()
}
