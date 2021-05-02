package main

import "fmt"

func main() {
	fmt.Println("slice-examples")
	// slices are basically arrays with size not pre-defined
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, 3) // here length is mentioned as 3 but still we added elements in y then length gets adjusted automatically.
	// make(1st,2nd,3rd) - 1st will be slice type and 2nd will be length and 3rd will be capacity
	// we can use make function to create a slice first parameter deines the type of slice
	//next argument defines the size of the slice
	copy(y, x) // x is copied to y but only first three because we mentioned only size of 3 but still we can add, we cant copy exceeding the limit mentioned
	// append function is used to add more elements into  the slice first argument defines which slice we should append
	// and remaining paramaters tells he elements that need to be added into the slice
	y = append(y, 6, 7, 8)
	fmt.Println(x, y)
	fmt.Println("capacity of y :", cap(y))
	fmt.Println("length of y :", len(y))
}
