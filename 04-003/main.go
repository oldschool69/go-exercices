package main

import "fmt"

func main() {

	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	s2 := s[:5]
	fmt.Println(s2)
	s3 := s[5:]
	fmt.Println(s3)
	s4 := s[2:7]
	fmt.Println(s4)
	s5 := s[1:6]
	fmt.Println(s5)

}
