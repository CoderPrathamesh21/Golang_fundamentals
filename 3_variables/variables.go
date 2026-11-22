package main

import "fmt"

func main() {
	var name string = "Prathamesh"
	fmt.Println("Hello, " + name + "! Welcome to Go programming.")

	var language = "Golang"
	fmt.Println("You are learning " + language + ".")

	//shothand syntax
	age := 25
	fmt.Println("Your age is:", age)

	//multiple variable declaration
	var x, y int = 5, 10
	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)

	//multiple variable declaration with shorthand syntax
	a, b := 15, 20
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)
}
