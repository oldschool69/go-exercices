package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}


func main() {

	s := square{
		side: 4,
	}

	c := circle{
		radius: 10.23,
	}

	info(s)
	info(c)
}

func info(s shape) {
	fmt.Println("The area of shape is: ", s.area())
}
