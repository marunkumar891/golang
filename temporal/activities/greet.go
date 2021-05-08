package activities

import (
	"context"
	"fmt"
)

func Greet(ctx context.Context, user string) error {

	fmt.Print("sendGreeting activity called")
	fmt.Printf("greeting sent to user : %v", user)
	return nil
}
