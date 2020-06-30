package main

import (
	"fmt"
)

func main() {
	// Cria um channel que recebera valores do tipo int
	c := make(chan int)

	// Passar o valor para o canal dentro da goroutine
	// Se executar da mesma thread é deadlock
	go func(){
		c <- 42
	}()

	//recupera valor do canal
	fmt.Println(<-c)

	//Fazer a mesma coisa acima usando buffered channel, sem goroutine
	c2 := make(chan int, 1)

	c2 <- 47

	fmt.Println(<-c2)
}
