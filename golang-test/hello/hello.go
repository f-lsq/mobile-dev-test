package main // declares a main package

import (
	"fmt"

	"rsc.io/quote"
) // imports the 'fmt' package for formatting text
// package includes printing to the console

func main() { // main function executes by default when you run the main package
	fmt.Println(quote.Go())
}
