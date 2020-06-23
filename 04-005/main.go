package main

import "fmt"

func main() {

	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(s)
	// Removendo os elementos 45 até 47 do slice
	s = append(s[:3], s[6:]...)
	fmt.Println(s)

}
