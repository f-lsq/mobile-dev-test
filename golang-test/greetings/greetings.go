package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// function with capital letter can be called by functions not in the same package

	// Return a greeting that embeds the name in a message.
	message:= fmt.Sprintf("Hi, %v. Welcome!", name)
	// ':=' is a shortcut for declaring and initialising a variable in one line (value on right is used to determine 'message' variable's type)
	// long way:
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}