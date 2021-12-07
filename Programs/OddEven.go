package main

import (
	"fmt"
)

func oddeven(val int) {
	if val%2 == 0 {
		fmt.Print("The value", val, "is Even")
	} else {
		fmt.Print("The value", val, "is Odd")
	}

}
func main() {
	oddeven(10)

}
