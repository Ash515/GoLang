package main

import (
	"fmt"
)

func main() {
	fmt.Print("Maps in Go")
	map_1 := map[int]string{
		5001: "Ashwin", 5002: "Joseph", 5003: "Kasim",
	}
	fmt.Println("The Map elements are:", map_1)

	for key, value := range map_1 {
		fmt.Print(key, value)
		fmt.Print("\n")
	}

}
