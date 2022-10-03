// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"os"
)

func main() {
	// TYPE YOUR CODE HERE
	var l int = len(os.Args)                    //len people
	fmt.Println("Hello my ", l-1, " friends :") //to no count the path file jaja

	for i := 1; i < l; i++ {
		fmt.Println("Hello great ", os.Args[i])
	}

	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.
}
