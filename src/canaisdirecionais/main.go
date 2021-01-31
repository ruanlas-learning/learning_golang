package main

import (
	"fmt"
)

func main() {
	canal := make(chan int) //declaração de canal bidirecional

	go send(canal)

	receive(canal)
}

func send(s chan<- int)  { //"uso" do canal de forma direcional (coloca dados no canal)
	s <- 42
}

func receive(r <-chan int)  { //"uso" do canal de forma direcional (retira dados do canal)
	fmt.Println("O valor recebido do canal foi:", <-r)
}