package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var output strings.Builder

	symbols := "@_."
	numeric := "0123456789"
	chars := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"

	domain := symbols + numeric + chars
	k := len(domain)
	fmt.Println(domain)

	var len int
	fmt.Println("please enter the length of your password")
	fmt.Scan(&len)

	for i := 0; i < len; i++ {
		index := rand.Intn(k)
		val := domain[index]
		output.WriteString(string(val))
	}

	fmt.Println("your password : ", output.String())
}
