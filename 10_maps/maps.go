package main

import (
	"fmt"
	"maps"
)

func main() {

	//creating a map
	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 3
	m["orange"] = 7

	//accessing map values
	fmt.Println(m["apple"])  // Output: 5
	fmt.Println(m["banana"]) // Output: 3
	fmt.Println(m["orange"]) // Output: 7
	fmt.Println(m["Guava"])  // Output: 0 (default value for int)

	fmt.Println("==============================")

	fmt.Println(len(m))

	fmt.Println("==============================")

	//checking if a key exists in the map
	value, ok := m["grape"]
	if ok {
		fmt.Println("Grape exists with value:", value)
	} else {
		fmt.Println("Grape does not exist in the map")
	}

	//deleting a key-value pair from the map
	delete(m, "banana")

	//iterating over the map
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

	fmt.Println("==============================")

	m1 := make(map[string]string)
	m1["apple"] = "red"
	m1["banana"] = "yellow"
	m1["orange"] = "orange"

	// iterating over the map
	for key, value := range m1 {
		fmt.Println(key, ":", value)
	}

	//clearing a map
	clear(m1)
	fmt.Println(m1)

	fmt.Println("==============================")

	m2 := map[string]int{"Rohit": 1, "Gill": 2, "Virat": 3, "Shreyas": 4, "Rahul": 5}
	fmt.Println(m2)

	fmt.Println("==============================")

	m3 := map[string]int{"India": 3, "Sri Lanka": 1, "Pakistan": 1, "Australia": 2}
	m4 := map[string]int{"India": 3, "Sri Lanka": 1, "Pakistan": 1, "Australia": 2}

	fmt.Println(maps.Equal(m3, m4))
}
