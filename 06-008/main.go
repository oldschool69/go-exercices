package main

import "fmt"

func main() {

	callBack := func(status string) {
		if status == "done" {
			fmt.Println("Done!")
		} else if status == "working" {
			fmt.Println("Still working...")
		}
	}

	fmt.Println("Starting...")
	foo(callBack, 10)

}

func foo(cb func(status string), number int) {

	for i := 1; i <= number; i++ {
		cb("working")
	}

	// Job done, notify by calling the callback
	cb("done")
}
