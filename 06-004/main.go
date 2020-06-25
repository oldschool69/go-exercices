package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Printf("My name is %v and I'm %v years old", p.first, p.age)
}

func main() {

	p := person{
		first: "Flavio",
		last: "Oliveira",
		age: 47,
	}

	p.speak()

}
