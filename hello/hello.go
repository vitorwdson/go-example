package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ") // Prefix the error output with "greetings: "
	log.SetFlags(0) // Disable printing the time and line number on errors

	name := "Vitor"
	message, err := greetings.Hello(name) // Extracting both the message and error

	if err != nil { // Checking if an error occurred
		log.Fatal(err) // Prints the error and exits the program
	}

	fmt.Println(message) // If no error happened, print the message
}