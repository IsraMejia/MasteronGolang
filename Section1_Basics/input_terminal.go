package main

import (
	"fmt"
	"os"
)

func main() {
	//go run input_terminal.go hi yo
	/*var Args []string
	fmt.Println("Input terminal is: ", Args)*/
	fmt.Printf("\n %#v\n", os.Args)
	fmt.Println("\nPath:", os.Args[0])
	fmt.Println("\nFirst Argument: ", os.Args[1])
	fmt.Println("\n2 Argument: ", os.Args[2])
	fmt.Println("\n3 Argument: ", os.Args[3])

	fmt.Println("\nNumber of items insade os.Args", len(os.Args))
}
