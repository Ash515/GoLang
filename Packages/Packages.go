/*
Packages are the most powerful part of the Go language.
The purpose of a package is to design and maintain a large number
of programs by grouping related features together into single units
so that they can be easy to maintain and understand and independent
of the other package programs. This modularity allows them to share
and reuse. In Go language, every package is defined with a different
name and that name is close to their functionality like “strings”
package and it contains methods and functions that only related to strings.
*/

// Go program to illustrate the
// concept of packages
// Package declaration
package main

// Importing multiple packages
import (
	"bytes"
	"fmt"
	"sort"
)

func main() {

	// Creating and initializing slice
	// Using shorthand declaration
	slice_1 := []byte{'*', 'G', 'e', 'e', 'k', 's', 'f',
		'o', 'r', 'G', 'e', 'e', 'k', 's', '^', '^'}
	slice_2 := []string{"Gee", "ks", "for", "Gee", "ks"}

	// Displaying slices
	fmt.Println("Original Slice:")
	fmt.Printf("Slice 1 : %s", slice_1)
	fmt.Println("\nSlice 2: ", slice_2)

	// Trimming specified leading
	// and trailing Unicode points
	// from the given slice of bytes
	// Using Trim function
	res := bytes.Trim(slice_1, "*^")
	fmt.Printf("\nNew Slice : %s", res)

	// Sorting slice 2
	// Using Strings function
	sort.Strings(slice_2)
	fmt.Println("\nSorted slice:", slice_2)
}
