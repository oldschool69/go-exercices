package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main(){

	var counter int64

	const numGoRoutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoRoutines)
	fmt.Println("numGoRoutines: ", numGoRoutines)
	for i := 0; i < numGoRoutines; i++ {
		go func(){
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter: ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("numGoRoutines: ", numGoRoutines)
	fmt.Println("counter: ", counter)


}