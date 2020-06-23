package main

import "fmt"

func main() {

	a := [5]int{}

	fmt.Println(a)

	a[0] = 25
	a[1] = 39
	a[2] = 105
	a[3] = 77
	a[4] = 1002

	fmt.Println(a)

	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", a)


}
