package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	//x = 45
	//y = "Flavio Oliveira"
	//z = false
	//vai imprimir zero values (0 "" false),
	//que são valores padrão atribuidos pelo compilador
	//quando nenhum valor foi atribuído para a variável
	//fmt.Println(x, y, z)

	s := fmt.Sprintf("x: %v y: %v z: %v", x, y, z)

	fmt.Println("s => ", s)



}
