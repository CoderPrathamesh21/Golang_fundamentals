package main

import "fmt"

// func printSlice(items []int) {

// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(languages []string) {

// 	for _, language := range languages {
// 		fmt.Println(language)
// 	}
// }

func printSlice[T comparable](items []T) { //[T any] can be replaced with [T interface{}] or [T int | string | bool etc.] or [T comparable]

	for _, item := range items {
		fmt.Println(item)
	}
}

// type stack struct {
// 	elements []int
// }

type stack[T any] struct {
	elements []T
}

func main() {

	// printSlice([]int{1, 2, 3})
	// fmt.Println("=====================")

	// languages := []string{"golang", "java", "python"}
	// printStringSlice(languages)
	languages := []string{"golang", "java", "python"}
	vals := []bool{true, false, true}
	printSlice(languages)

	fmt.Println("=====================")
	printSlice(vals)

	fmt.Println("=====================")
	myStack := stack[string]{
		elements: []string{"golang", "java", "python"},
	}
	fmt.Println(myStack)

}
