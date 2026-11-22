package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {

	fmt.Println(1, 2, true, "golang", 3.14)

	nums := []int{03, 12, 21, 30}
	result := sum(nums...)
	fmt.Println(result)

}
