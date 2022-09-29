package main

import "fmt"

func main() {
	/*var speed int
	fmt.Println("Speed = ", speed)

	//Declaration
	var (
		video    string = "Salu2"
		actor    string = "Isra"
		duration int
		current  int
	) //when we will use a related var we can declaration like this

	fmt.Println("video, duration, current , actor= ", video, duration, current, actor)
	//fmt.Println(video, duration, current)
	actor2 := "Isra2"
	fmt.Println("actor2:", actor2)
	*/
	speed := 100 // int
	force := 2.5 //float64

	//speed = speed * force  //error Types  variables
	speed = int(float64(speed) * force) // cast to the same type to solve, speed is a int
	fmt.Println("speed = ", speed)

}
