package main

import "fmt"

func add(num1 int, num2 int) int {
	// arguments should b  in this format name of variable and then the type
	// at last return type of that function should be mentioned
	var result = num1 + num2
	return result
}
func main() {
	num1 := 5
	num2 := 7
	res := add(num1, num2)
	fmt.Println(res)

}
