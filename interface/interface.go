package main

import (
	"fmt"
)

type person interface {
	speak() // return type of interface is always not necesary
}

//the common function implemented by both class is added in interface like a common feature
//and always method defined in an interface should be implemented

type student struct {
	name string
}
type teacher struct {
	name string
}

func (obj *student) speak() string {
	return fmt.Sprintf("hello i am a student and my name is : %s \n", obj.name)
}

func (obj *teacher) speak() string {
	return fmt.Sprintf("hello i am a teacher and my name is : %s \n", obj.name)
}
func add(x interface{}) {
	fmt.Println("the value is:", x)
	fmt.Println()
} // empty interface can be used like this also, when data type of parameter is not known
func main() {
	fmt.Println("Go-Interface Examples")
	stud := &student{"arunkumar"}
	tech := &teacher{"mak"}
	fmt.Println(stud.speak())
	fmt.Println(tech.speak())

	var x = 10
	add(x)
}
