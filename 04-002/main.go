package main

import "fmt"

func main() {

	s := []int{20, 1, 13, 6, 4545, 88, 989, 10, 2, 44}

	fmt.Println(s)

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", s)


}
