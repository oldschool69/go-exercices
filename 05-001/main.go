package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {

	p1 := person{
		first: "Flavio",
		last:  "Oliveira",
		flavors: []string{
			"baunilha",
			"passas ao rum",
			"chocolate",
		},
	}

	p2 := person{
		first: "Pablo",
		last:  "Escobar",
		flavors: []string{
			"pistache",
			"kiwui",
		},
	}

	fmt.Printf("p1 %v %v \n", p1.first, p1.last)
	for _, v := range p1.flavors {
		fmt.Printf("\t%v\n", v)
	}
	fmt.Printf("p2 %v %v \n", p2.first, p2.last)
	for _, v := range p2.flavors {
		fmt.Printf("\t%v\n", v)
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Printf("Name %v %v\n", v.first, v.last)
		for i, flavor := range v.flavors {
			fmt.Printf("\tflavor %v %v\n", i, flavor)
		}
	}

}
