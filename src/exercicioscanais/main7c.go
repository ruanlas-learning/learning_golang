package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	q := make(chan bool)
	goroutines := 10

	go send(goroutines, c, q)
	receive(c, q)

	fmt.Println("exit")
}

func send(funcoes int, canal chan<- string, quit chan<- bool) {
	var wg sync.WaitGroup
	wg.Add(funcoes)
	for i := 0; i < funcoes; i++ {
		go func(x int) {

			for y := 0; y < 10; y++ {
				canal <- fmt.Sprintf("Goroutine %v lança nº [%v]", x, y)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	quit <- true
}

func receive(c <-chan string, q <-chan bool) {
	for{
		select {
			case v := <-c:
				fmt.Println(v)
			case <-q:
				return
		}
	}
	
}