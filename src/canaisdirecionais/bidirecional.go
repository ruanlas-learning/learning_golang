package main

import (
	"fmt"
)

func main() {
	c := make(chan int) //declaração de canal bidirecional
	cr := make(<-chan int) //declaração de canal unidirecional "receive"
	cs := make(chan<- int) //declaração de canal unidirecional "send"

	fmt.Println("--------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// general to specific converts
	fmt.Println("--------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	// não funciona!!!
	// cs = cr
	// cs = (chan<- int)(cr)
	
	// Funciona!!!
	// cr = (<-chan int)(c)
}