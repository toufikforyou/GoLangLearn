package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/toufikforyou/golanglearn/auth"
	"github.com/toufikforyou/golanglearn/user"
)

func main() {
	auth.LoginWithCredentials("toufikforyou", "secret")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		// Name:  "John Doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)

}
