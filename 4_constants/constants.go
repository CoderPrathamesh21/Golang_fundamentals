package main

import "fmt"

const age = 27

func main() {

	const name = "Prathamesh"
	fmt.Println("Hello, " + name + "! Welcome to Go programming.")

	fmt.Println("My age is:", age)

	const (
		country = "India"
		state   = "Maharashtra"
		city    = "Mumbai"
	)
	fmt.Println("You live in " + city + ", " + state + ", " + country + ".")

	const (
		host = "localhost"
		port = 8080
	)
	fmt.Println("Server is running on " + host + ":" + fmt.Sprint(port))

}
