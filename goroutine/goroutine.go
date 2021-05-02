package main

import (
	"fmt"
	"time"
)

func speak(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 1000)
	}
}
func main() {
	fmt.Println("Simple GoRoutine Example.!")
	go speak("cat") // using go will make this as a goroutine and this function call will run in background
	go speak("dog") // when you make both functions as go routines then use scanln to run the program or else the main function will terminate.

	fmt.Scanln() // using this will make both goroutines execute untill user gives some input.(only when its a inifinte loop in the function)
}
