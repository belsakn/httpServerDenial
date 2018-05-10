package main

import (
	"fmt"
)

func main() {

	printUsage()

}

func printUsage() {
	fmt.Printf("\n Usage: \n\n \t interview-mali8 command [arguments]")
	fmt.Printf("\n\n The commands are: \n")
	fmt.Printf("\n \t server \t run HTTP Denial-of-Service protection system on port 8080")
	fmt.Printf("\n \t client \t run HTTP client")
	fmt.Printf("\n \t help \t\t print usage instructions")
	fmt.Printf("\n")
}
