package main

import (
	"fmt"
)

func main() {

	birth := 1973
	current := 2020

	for {
		fmt.Println(birth)
		if birth == current {
			break
		}
		birth++
	}
}
