package main

import "fmt"

func main() {

	//while loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("================================")

	//infinite loop
	// for {
	// 	fmt.Println("This is an infinite loop")
	// }

	//classic for loop
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	fmt.Println("================================")

	//break usage
	for k := 1; k <= 10; k++ {
		if k == 6 {
			break
		}
		fmt.Println(k)
	}

	fmt.Println("================================")

	//usage continue
	for l := 1; l <= 10; l++ {
		if l%2 == 0 {
			continue
		}
		fmt.Println(l)
	}

	fmt.Println("================================")

	for p := range 4 {
		fmt.Println(p)
	}

}
