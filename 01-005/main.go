package main

import "fmt"

type hotdog int
var x hotdog
var y int

func main ()  {
	fmt.Println("x: ", x)
	fmt.Printf("x type: %T\n", x)
	x = 666
	fmt.Println("x: ", x)
	fmt.Printf("y type before: %T\n", y)
	y := int(x)
	fmt.Println("y: ", y)
	fmt.Printf("y type after: %T\n", y)
}
