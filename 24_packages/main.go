package main

import (
	"fmt"

	"github.com/coderprathamesh21/podcast/auth"
	"github.com/coderprathamesh21/podcast/user"
	"github.com/fatih/color"
)

func main() {

	auth.LoginWithCredentials("coderprathamesh", "secret")

	session := auth.GetSession()
	fmt.Println("Session", session)

	user := user.User{
		Name:  "David",
		Email: "dave@email.com",
	}

	fmt.Println(user.Name, user.Email)
	color.Red(user.Name, user.Email)
}
