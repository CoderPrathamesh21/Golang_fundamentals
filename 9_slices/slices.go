package main

import (
	"fmt"
	"slices"
)

func main() {

	var nums []int

	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	fmt.Println("==============================")

	//var vals = make([]int, 3, 5)
	var vals = make([]int, 0, 5)

	vals = append(vals, 10)
	vals = append(vals, 20)
	vals = append(vals, 30)

	fmt.Println(vals)
	fmt.Println(vals == nil)
	fmt.Println(cap(vals))

	fmt.Println("==============================")

	vals2 := []int{2}
	fmt.Println(vals2)
	fmt.Println(vals2 == nil)
	fmt.Println(cap(vals2))

	vals2[0] = 100
	fmt.Println(vals2)

	fmt.Println("==============================")

	var numbers = make([]int, 3, 4)
	numbers = append(numbers, 1)

	var numbers2 = make([]int, len(numbers))
	fmt.Println(numbers, numbers2)

	//copying slices
	copy(numbers2, numbers)
	fmt.Println(numbers2, numbers)

	fmt.Println("==============================")

	//slicing a slice
	var nums2 = []int{11, 21, 32, 44, 55}
	var sliced = nums2[1:4]
	fmt.Println(sliced)

	fmt.Println("==============================")

	//slicing a slice with only start index
	var sliced2 = nums2[2:]
	fmt.Println(sliced2)

	fmt.Println("==============================")

	//slicing a slice with only end index
	var sliced3 = nums2[:3]
	fmt.Println(sliced3)

	fmt.Println("==============================")

	//slicing a slice with start and end index
	var sliced4 = nums2[1:4]
	fmt.Println(sliced4)

	fmt.Println("==============================")

	var n1 = []int{1, 2, 3}
	var n2 = []int{1, 2, 3}

	fmt.Println(slices.Equal(n1, n2))

	var values = [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(values)
	fmt.Println("==============================")
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println("==============================")
	fmt.Println(values[0][1])
	fmt.Println(values[1][0])
	fmt.Println(values[2][1])
	fmt.Println("==============================")
	fmt.Println(slices.Equal(values[0], values[1]))
	fmt.Println(slices.Equal(values[0], values[0]))

	fmt.Println("==============================")

	var n3 = []int{1, 2, 3}
	var n4 = []int{1, 2, 3, 4}

	fmt.Println(slices.Equal(n3, n4))

}
