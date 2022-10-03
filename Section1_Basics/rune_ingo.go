// Simple Go program to illustrate
// how to create a rune
// this code is from https://www.geeksforgeeks.org/rune-in-golang/ , I saved
// because i like it the example
package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Creating a rune
	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	// Displaying rune and its type
	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s", rune1,
		rune1, reflect.TypeOf(rune1))

	fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %s", rune2,
		rune2, reflect.TypeOf(rune2))

	fmt.Printf("\nRune 3: Unicode: %U; Type: %s", rune3,
		reflect.TypeOf(rune3))

}
