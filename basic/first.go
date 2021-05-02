// package main something like the mandatory package that needs
// to be added in all go programs.
package main

// fmt is a module that helps to diplay messages in output
// field and also provides various other features.
import "fmt"

// this will be entry point for execuring the program
func main() {
	fmt.Println("hello,go...") // only a single character can be declared within single quotation.
}

// instead of go run if we use go build then an exe file will be created
// that can be shared very easily.
