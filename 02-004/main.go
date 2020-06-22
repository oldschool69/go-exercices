package main

import "fmt"

func main() {

	x := 10
	fmt.Printf("%d\t\t%b\t\t%#X\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\t\t%#X", y, y, y)

}
