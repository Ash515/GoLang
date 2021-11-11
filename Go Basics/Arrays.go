package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3, 4}
	arr2 := [...]float32{1.2, 2.3, 3.4}
	arr3 := [4]int{10, 1: 20, 2: 30}
	fmt.Print(arr1[0])
	fmt.Println(arr2)
	fmt.Print(len(arr1))
	fmt.Print(arr3)
}
