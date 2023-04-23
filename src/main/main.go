package main

import (
	"404coder.com/greetings"
	"fmt"
	"log"
)

import "rsc.io/quote"

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())

	names := []string{"Strawberry", "Darwin", "Cao", "Zephyr"}
	// Get a greeting messages and print it.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	// If no error was returned, print the returned messages
	// to the console.
	for _, name := range names {
		fmt.Println(messages[name])
	}
}
