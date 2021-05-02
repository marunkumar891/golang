package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name     string
	Password string
	LoggedIn time.Time
	// since these variables are used in json encoding first letter shoul be capital
	// if we need to use any other formats we must map the json object with the struct tags
	/*
		Name     string		`json:"name"`
		Password string     `json:"password"` // if we use - instead of passsword that means password field is hidden and it will not be displayed in the output
		LoggedIn time.Time	`json:"createdAt"`
	*/
}

// To encode JSON data we use the Marshal function.
// basic representation : func Marshal(v interface{}) ([]byte, error)
func marshal(u1 *User) []byte {
	res, err := json.MarshalIndent(u1, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return res

}

// To decode JSON data we use the Unmarshal function.
// func Unmarshal(data []byte, v interface{}) error

func unmarshal(u1 *User, res []byte) string {
	err := json.Unmarshal(res, u1)
	if err != nil {
		return fmt.Sprintf("%s", (err))
	}
	return fmt.Sprintf("Name = %s, Password = %s, LogIn = %s", u1.Name, u1.Password, u1.LoggedIn)
}

func main() {
	fmt.Println("program started at:", time.Now())
	u1 := &User{"arunkumar", "qwerty123", time.Now()}
	//fmt.Println(u1)
	res := marshal(u1)
	fmt.Println(string(res))
	fmt.Println(unmarshal(u1, res))

}
