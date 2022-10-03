// ---------------------------------------------------------
// EXERCISE: Discard The File
//
//  1. Print only the directory using `path.Split`
//
//  2. Discard the file part
//
// RESTRICTION
//  Use short declaration
//
// EXPECTED OUTPUT
//  secret/
// ---------------------------------------------------------

package main

import (
	"fmt"
	"path"
)

func main() {
	dir, _ := path.Split("secret/file.txt")
	//dir, fName := path.Split("secret/file.txt")

	fmt.Println(dir)
	//fmt.Println(fName)
}
