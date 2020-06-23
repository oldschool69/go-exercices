package main

import "fmt"

func main() {

	s := struct {
		name string
		age int
		musicStyles []string
		addresses map[string]string
	}{
		name: "Flavio",
		age: 42,
		musicStyles: []string{
			"Blues",
			"Jazz",
			"Rock and Roll",
		},
		addresses: map[string]string{
			"home":"Rua Jose Luiz Camargo Moreira, 33",
			"work":"Rua Aguacú, 171",
		},
	}

	fmt.Println(s)

	for i, v := range s.musicStyles {
		fmt.Println(i, v)
	}

	fmt.Println("Home address: ", s.addresses["home"])
}
