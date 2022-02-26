package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Vitor")
	fmt.Println(message)
}