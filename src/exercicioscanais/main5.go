package main

import (
	"fmt"
)

func main() {
	// c := make(chan int, 1)
	c := make(chan int)

	// c <- 6
	go func() {
		c <- 5
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)

}