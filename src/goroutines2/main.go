package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	novasgoroutines(100)
	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}

func novasgoroutines(i int)  {
	wg.Add(i)

	for j := 0; j < i; j++ {
		x := j

		go func(i int){
			fmt.Println("Eu sou a goroutine número:", i+1)
			wg.Done()
		}(x)
	}
}