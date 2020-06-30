package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Processor arch: ", runtime.GOARCH)
	var counter int64
	const numGoRoutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("processing counter: ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter final value: ", counter)
	fmt.Println("Exiting main func")
}
