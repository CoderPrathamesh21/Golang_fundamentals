package main

import "time"

func main() {

	// i := 15

	// switch i {
	// case 1:
	// 	println("One")
	// case 2:
	// 	println("Two")
	// case 3:
	// 	println("Three")
	// case 4:
	// 	println("Four")
	// case 5:
	// 	println("Five")
	// default:
	// 	println("Unknown number")
	// }

	//multiple conditions in a single case

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend!")
	default:
		println("It's a weekday.")
	}

	//type switch

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("I'm an integer")
		case string:
			println("I'm a string")
		case float64:
			println("I'm a float64")
		case bool:
			println("I'm a boolean")
		default:
			println("Unknown type")
		}
	}

	whoAmI(42)
	whoAmI("Hello, Go!")
	whoAmI(3.14)
	whoAmI(true)
	whoAmI(nil)
}
