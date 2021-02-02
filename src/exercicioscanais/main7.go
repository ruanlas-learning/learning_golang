package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	goroutines := 10

	send(goroutines, c)
	receive(c)

	fmt.Println("exit")
}

func send(funcoes int, canal chan<- string) {

	for i := 0; i < funcoes; i++ {
		go func(x int) {

			for y := 0; y < 10; y++ {
				canal <- fmt.Sprintf("Goroutine %v lança nº [%v]", x, y)
			}

		}(i)
	}
}

func receive(c <-chan string) {
	// for v := range c {
	// 	fmt.Println(v)
	// }

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}