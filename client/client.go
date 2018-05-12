package client

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Client() {

	setupCloseHandler()

	startClient(fmt.Sprintf("http://localhost:8080/?clientId=1"))

}

func startClient(url string) {
	for {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		log.Printf("Get %s response code is %d", url, resp.StatusCode)
		r := rand.Intn(1222)
		time.Sleep(time.Duration(r) * time.Millisecond)
	}
}

func setupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("# Ctrl+C pressed")
		os.Exit(0)
	}()
}
