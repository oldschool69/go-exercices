package main

import "fmt"

func main() {

	fmt.Println("main begins")
	defer foo()
	fmt.Println("main ends")

}

func foo() {
	fmt.Println("foo begins")
	// Do some work here
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("foo ends")
}
