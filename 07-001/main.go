package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	p := person{
		name: "Flavio",
		age: 47,
	}
	fmt.Println("p before: ", p)
	changeMe(&p)
	fmt.Println("p after: ", p)
}

func changeMe(p *person) {
	// Para estruturas é possível usar
	// p.name ou (*p).name, para dereferenciar o campo de uma struct
	p.name = "Josias"
	(*p).age = 54
}
