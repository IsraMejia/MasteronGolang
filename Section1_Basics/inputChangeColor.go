// when you run this code you have to use the next comand
// go run .\inputChangeColor.go blue
package main

import (
	"fmt"
	"os"
)

func main() {
	var color string = "red"
	//color = os.Args[0] // error this is the path
	color = os.Args[1] // here i have my first passed argument

	fmt.Println("\nNow color is = ", color)

}
