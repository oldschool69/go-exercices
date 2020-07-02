package main

import "fmt"

type customErr struct {
	info string
}

func main() {

	ce := customErr{
		info: "critical error",
	}

	foo(ce)

}

func (ce customErr) Error() string {
	return fmt.Sprintf("A custom error occurred: %v", ce.info)
}

func foo(err error) {
	fmt.Println(err.Error())
	//Para acessar o campo info da estrutura customError recebida
	//aqui, é necessário utilizar assertion
	fmt.Println(err.(customErr).info)
}
