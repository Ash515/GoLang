package main

import "fmt"

func main() {

	// Creating and initializing empty map
	// Using var keyword
	var map_1 map[int]int

	// Checking if the map is nil or not
	if map_1 == nil {

		fmt.Println("True")
	} else {

		fmt.Println("False")
	}

	map_2 := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println("Map-2: ", map_2)
}
