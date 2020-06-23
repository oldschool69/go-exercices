package main

import "fmt"

func main() {
	x := [][]string{
		{"James","Bond","Shaken","not stirred"},
		{"Miss","Moneypenny","Helloooo","James"},
	}

	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
		for j, vv := range v {
			fmt.Println(j, vv)
		}
	}
}
