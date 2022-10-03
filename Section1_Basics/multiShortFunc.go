// EXERCISE: Multi Short Func
//
// 	1. Declare two variables using multiple short declaration syntax
//
//  2. Initialize the variables using `multi` function below
//
//  3. Discard the 1st variable's value in the declaration
//
//  4. Print only the 2nd variable

package main

import (
	"fmt"
)

func main() {
	_, b, _ := multi()

	fmt.Println(b)
	//fmt.Println(_) // error, we discard the 1st value with '_'
	// _ = define and use the unused variable using Blank Identifier
}

func multi() (int, int, int) { //(int, int, int) return type values
	return 5, 4, 6
}
