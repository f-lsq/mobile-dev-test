package main // declares a main package
// NOTE: Code executed as an application must be in the main package

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
) // imports the 'fmt' package for formatting text
// package includes printing to the console

func main() { // main function executes by default when you run the main package
	fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including the log entry prefix and
	// a flag to disable printing the time, source file and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message and print it
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console
	fmt.Println(message)
}
