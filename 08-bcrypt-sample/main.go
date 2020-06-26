package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"

	// Before using bcrypt, need to import package by running the command:
	// go get golang.org/x/crypto/bcrypt
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	//fmt.Println(s)
	//fmt.Println(bs)

	loginPassword := "password123"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("Invalid password, can't login: ")
		return
	}
	fmt.Println("You can login!")


}
