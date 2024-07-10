package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// function with capital letter can be called by functions not in the same package -> exported name

	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	//If a name was received, return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// ':=' is a shortcut for declaring and initialising a variable in one line (value on right is used to determine 'message' variable's type)
	// long way:
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil // 'nil' means no error
}
