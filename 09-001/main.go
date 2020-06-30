package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var numGoRoutines = 2
	var waitGroup sync.WaitGroup
	waitGroup.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		go func(num int) {
			fmt.Println("Starting go routine: ", num)
			time.Sleep(time.Second)
			fmt.Println()
			fmt.Println("Finishing go routine: ", num)
			waitGroup.Done()
		}(i + 1)
	}

	waitGroup.Wait()
	fmt.Println("Exiting main goroutine")

}
