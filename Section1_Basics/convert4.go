// ---------------------------------------------------------
// EXERCISE: Convert and Fix #4
//
//  Fix the code.
// age := 2
// fmt.Println(int(7.5) + int(age))
//
// EXPECTED OUTPUT
//  9.5
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	age := 2
	fmt.Println(7.5 + float64(age))
}
