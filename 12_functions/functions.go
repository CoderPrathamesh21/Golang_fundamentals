package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "java", "python"
}

// function calling function
func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {

	sum := add(21, 26)
	fmt.Println("Sum: ", sum)

	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)

	fn := processIt()
	fmt.Println(fn(6))
}
