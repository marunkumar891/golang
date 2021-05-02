package main

import (
	"fmt"
)

type Person struct {
	fname, lname string
	age          int
}

func printPerson(p *Person) {
	fmt.Printf("person name is %s and age is %d", p.fname+p.lname, p.age)
}
func (p *Person) changeAge() {
	p.age = 21
	fmt.Println("\nAfter modifying the age:", p.age)
}
func (p *Person) String() string {
	return fmt.Sprintf("name is %s", p.fname+p.lname)
}
func main() {
	fmt.Println("Stuct concepts : ")
	p := &Person{"arun", "kumar", 20}
	printPerson(p)
	// This function type is used to perform any operation on person data mambers
	// this is similar to member function related to person struct
	p.changeAge()
	// getting a string output from a fucntion
	fmt.Println(p)
	// On the User type, we also define the String() method required by the fmt.Stringer interface.
	// This will be called automatically when we pass the type to fmt.Println and gives us a chance to
	// produce a nicely formatted version of our struct.
}
