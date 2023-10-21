package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set the properties of the predefined Logger, including
	// the log entry prefix and flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Denes")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	// A slice of names.
	names := []string{"Pati", "Morzsi", "Vitez"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}
