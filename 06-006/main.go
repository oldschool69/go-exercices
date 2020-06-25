package main

import "fmt"

func main() {

	// Função anônima
	func(name string){
		fmt.Println("My name is ", name)
	}("Flavio")

	// atribuir função para variável e executar
	f := func(v1, v2 int) int{
		return v1 + v2
	}

	fmt.Println(f(30, 70))

	fmt.Println("That's it!")


}
