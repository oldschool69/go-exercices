package main

import "fmt"

type person struct {
	name string
	age int
}

func (p person) speak() {
	fmt.Println("My name is ", p.name)
}

type human interface {
	speak()
}

func main() {
	p := person{
		name: "Flavio",
		age: 47,
	}

	saySomething(&p)
	saySomething(p)
}

func saySomething(h human) {
	h.speak()
}
