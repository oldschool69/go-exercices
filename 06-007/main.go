package main

import "fmt"

func main() {
	f1 := foo([]int{1, 2, 3})
	f2 := foo([]int{4, 5, 6})
	fmt.Println(f1())
	fmt.Println(f2())
}



func foo(values []int) func() int {
	var sum int
	return func() int {
		for _, v := range values {
			sum += v
		}
		return sum
	}
}
