package main

import "fmt"

func main() {

	age := 11

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	var role = "admin"
	var hasAccess = true

	if role == "admin" && hasAccess {
		fmt.Println("You have admin access.")
	} else if role == "user" && hasAccess {
		fmt.Println("You have user access.")
	} else {
		fmt.Println("You do not have access.")
	}

	if voter_age := 20; voter_age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}
}
