package main

import (
	"fmt"
	// "sync"
)

func main() {
	c := make(chan string)
	goroutines := 10

	for i := 0; i < goroutines; i++ {
		go func(x int) {
			for y := 0; y < 10; y++ {
				c <- fmt.Sprintf("Goroutine %v lança nº [%v]", x, y)
			}
		}(i)
	}
	// close(c)
	receive(c)

	fmt.Println("exit")
}

func receive(c <-chan string) {
	// for v := range c {
	// 	fmt.Println(v)
	// }


	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}