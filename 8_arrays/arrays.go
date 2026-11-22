package main

import "fmt"

func main() {

	var nums [5]int
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50

	fmt.Println(len(nums))
	fmt.Println(nums[3])
	fmt.Println(cap(nums)) //capacity of array
	fmt.Println(nums)

	for i := range len(nums) {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()

	fmt.Println("==============================")
	var vals [4]bool
	vals[1] = true
	fmt.Println(vals)

	fmt.Println("==============================")
	//array initialization
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)

	fmt.Println("==============================")
	//array initialization with shorthand syntax
	arr2 := [4]string{"Go", "Python", "Java", "C++"}
	fmt.Println(arr2)

	fmt.Println("==============================")
	//multidimensional array
	var multiArr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(multiArr)

	fmt.Println("==============================")
}
