package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps or also called Dictionary")
	dict := make(map[int]string)
	dict[1] = "golang"
	dict[2] = "web dev"
	dict[3] = "back end"
	fmt.Println("Dictionary is:\n", dict)
	// we can iterate using for like this :
	for k, v := range dict {
		fmt.Printf("key : %d and value : %s\n", k, v)
	}
	val, stat := dict[2]
	//when we access the dictionary value it returns 2 thing 1-the value and 2-bool value true if exist else false
	fmt.Println(val, stat)
}
