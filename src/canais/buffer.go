package main

import (
	"fmt"
)

func main() {
	// configurando um buffer
	canal := make(chan int, 1) // o parametro 1 está especificando o numero de buffer
	canal <- 43
	fmt.Println(<-canal)
}