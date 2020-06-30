package main

import "fmt"

func main() {
	c := make(chan int)
	put(c)
	pull(c)

	fmt.Println("exiting main")
}

func pull(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func put(c chan<- int) {
	go func() {
		for i:=0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}