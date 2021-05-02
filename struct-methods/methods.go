package main

import (
	"fmt"
)

type Student struct {
	name   string
	grades []int
	age    int
}

func sum(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}
func (stud *Student) getage() {
	fmt.Printf("age of student is : %d", stud.age)
	fmt.Println()
}
func (stud *Student) gettotal() int {
	return sum(stud.grades)
}
func (stud *Student) checkresult() string {
	marks := stud.gettotal()
	if marks > 75 {
		return fmt.Sprintf("marks obtained : %d result passed", marks)
	} else {
		return fmt.Sprintf("marks obtained : %d result : withheld", marks)
	}

}
func main() {
	stud1 := &Student{"arunkuamr", []int{49, 37, 48}, 20}
	stud1.getage()
	fmt.Println("Total marks obtained : \n", stud1.gettotal())
	fmt.Println("Exam Result \n", stud1.checkresult())
}
