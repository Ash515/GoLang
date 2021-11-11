package main

import (
	"fmt"
)

func main() {
	/*
		Variable Declaration
	*/
	var studentname string
	var studentcgpa float32
	var a int = 10 // Traditional type of variable creation
	b := 20        //another type of variable creation compiler automatically detects the datatype
	c := a + b
	fmt.Print(c)
	fmt.Print(studentcgpa)
	fmt.Print(studentname)
}
