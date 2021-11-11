package main

import (
	"fmt"
)

func add(a int, b int) int {
	var c = a + b
	return c
}
func main() {
	var a, b, c int
	fmt.Println("Enter 1st Number:")
	fmt.Scan(&a)
	fmt.Println("Enter 2st Number:")
	fmt.Scan(&b)
	c = add(a, b)
	fmt.Println("Sum:%d\n", c)

}
