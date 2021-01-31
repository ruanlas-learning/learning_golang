package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	quit := make(chan int)

	go recebeQuit(canal, quit)
	enviaParaCanal(canal, quit)
}

func recebeQuit(canal, quit chan int){
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido:", <-canal)
	}
	quit <- 0
}

func enviaParaCanal(canal, quit chan int) {
	qualquerCoisa := 10
	for {
		select {
			case canal <- qualquerCoisa:
				qualquerCoisa++
			case <-quit:
				return
		}
	}
}