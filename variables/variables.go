// once after a variable is declared with specific data type
// its not possible to re decalre the same variable with other data type

package main
import "fmt"

// if a variable is declared and if its not used in the program then 
// compiler throws an error

func main() {

	const num = 9 //method to declare a constant value

	var name string = "name is a variable with string data type"
	fmt.Println(name)
	fmt.Println("this is a constant:",num)
}