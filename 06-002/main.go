package main

import (
	"fmt"
)

func main() {

	x := []int{1,2,3,4}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(values ...int) int {
	var sum int
	for _, v := range values {
		sum += v
	}
	return sum
}

func bar(values []int) int {

	var sum int
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum
}


