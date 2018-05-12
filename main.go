package main

import (
	"fmt"
	"interview-mali8/client"
	"interview-mali8/server"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		switch ok := os.Args[1]; ok {
		case "server":
			server.Server()
		case "client":
			client.Client()
		case "help":
			printUsage()
		default:
			printUsage()
		}
	} else {
		printUsage()
	}

}

func printUsage() {
	fmt.Printf("\n Usage: \n\n \t interview-mali8 command [arguments]")
	fmt.Printf("\n\n The commands are: \n")
	fmt.Printf("\n \t server \t\t run HTTP Denial-of-Service protection system on port 8080")
	fmt.Printf("\n \t client [clientNum] \t run clientNum HTTP clients")
	fmt.Printf("\n \t help \t\t\t print usage instructions")
	fmt.Printf("\n")
}
