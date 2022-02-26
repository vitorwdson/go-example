package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

// The init function will always run on the program startup
func init() {
	// Setting the seed for the rand module
	rand.Seed(time.Now().UnixNano())
}

func randomGreeting() string {
	// Creating a slice (an dynamically sized array) of greeting formats
	greetings := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return an random index of the slice
	return greetings[rand.Intn(len(greetings))]
}