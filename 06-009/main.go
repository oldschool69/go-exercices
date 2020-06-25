package main

import "fmt"

func main() {
	v1 := []int{1, 2, 3, 4}
	v2 := []int{10, 20, 30, 40}
	f1 := sum(v1...)
	f2 := sum(v2...)
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f2())
}


func sum(values ...int) func() int {
	var total int
	return func() int{
		for _, v := range values {
			total += v
		}
		return total
	}
}
