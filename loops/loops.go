package main

import "fmt"

func main() {
	for {
		if i := 3; i%10 == 0 {
			// if this if condition fails the this loop is going to be infinite
			break
		} else {
			fmt.Println(i)
		}
	}
}
