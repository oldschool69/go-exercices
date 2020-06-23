package main

import "fmt"

func main() {

	m := map[string][]string{}

	m["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Woman"}
	m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	m["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}
	m["flavio_oliveira"] = []string{"guitar", "beer", "woman"}

	// remove o valor do map
	delete(m, "bond_james")

	for k, v := range m {
		fmt.Println(k)
		for i, vv := range v {
			fmt.Printf("\t%v\t%v\n", i, vv)
		}
	}
}
