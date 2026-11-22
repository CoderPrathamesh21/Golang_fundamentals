package main

import (
	"fmt"
)

func main() {

	teams := []string{"CSK", "DC", "LSG", "MI", "PBKS"}

	//using for loop
	for i := 0; i < len(teams); i++ {
		fmt.Println(teams[i])
	}

	fmt.Println("==============================")

	//using range
	for i, team := range teams {
		fmt.Println(team, i)
	}

	fmt.Println("==============================")

	nums := []int{21, 26, 24}
	sum := 0
	//using range
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println("Sum: ", sum)

	fmt.Println("==============================")

	m := map[string]string{"Fname": " Prathamesh", "Lname": " Chikane"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("==============================")

	for i, c := range "golang" {
		//fmt.Println(i, c) // unicode values
		fmt.Println(i, string(c))
	}
}
