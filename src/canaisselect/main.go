package main

import (
	"fmt"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	x := 500

	go func (y int) {
		for i := 0; i < y; i++ {
			a <- i
		}
	}(x/2)

	go func (y int) {
		for i := 0; i < y; i++ {
			b <- i
		}
	}(x/2)

	for i := 0; i < x; i++ {
		// pode receber informações de vários canais
		select {
			case v := <-a:
				fmt.Println("Canal A:", v)
			case v := <-b:
				fmt.Println("Canal B:", v)
		}
	}
}