package main

import "fmt"

// pass by value
func changeNum(num int) {
	num = 6
	fmt.Println("InchangeNum", num)
}

// pass by reference which is pointer
func changeNumber(num *int) {
	*num = 6
	fmt.Println("InchangeNum", *num)
}

func main() {

	num := 1

	changeNum(num)
	fmt.Println("After changeNum in main", num)

	changeNumber(&num)
	fmt.Println("After changeNum in main", num)

}
