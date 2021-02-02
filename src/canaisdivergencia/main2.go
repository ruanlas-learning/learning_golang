package main

import (
	"fmt"
	// "math/rand"
	"time"
	"sync"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	// Trabalha de 5 em 5
	funcoes := 5

	go manda(100, canal1)
	go outra(funcoes, canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(funcoes int, canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for i := 0; i < funcoes; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)	
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000) //time.Duration(rand.Intn(1e3)))
	return n * 100
}