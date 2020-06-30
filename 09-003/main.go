package main

import (
	"fmt"
	"sync"
)

func main() {

	counter := 0
	const numGoRoutines = 100
	var wg sync.WaitGroup
	var mtx sync.Mutex
	wg.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		go func() {
			mtx.Lock()
			v := counter
			//runtime.Gosched()
			v++
			counter = v
			fmt.Println("processing counter: ", counter)
			mtx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter final value: ", counter)
	fmt.Println("Exiting main func")
}
