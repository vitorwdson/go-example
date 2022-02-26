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

func Hellos(names []string) (map[string]string, error) {
	// Initiates the messages map (types are specified as map[key]value)
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		// Creates a new key-value pair for the the name and message
		messages[name] = message
	}

	return messages, nil
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