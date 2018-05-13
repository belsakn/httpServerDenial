package main

import (
	"fmt"
	"interview-mali8/client"
	"interview-mali8/server"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) > 1 {
		switch ok := os.Args[1]; ok {
		case "server":
			server.Server()
		case "client":
			if len(os.Args) < 3 {
				fmt.Printf("\n Provide client argument! \n")
				printUsage()
				break
			}

			numberOfClients, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Printf("\n Provide client argument as whole number! \n")
				printUsage()
				break
			}

			client.Client(numberOfClients)
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
	fmt.Printf("\n \t client [clientNum] \t run clientNum of HTTP clients")
	fmt.Printf("\n \t help \t\t\t print usage instructions")
	fmt.Printf("\n")
}
