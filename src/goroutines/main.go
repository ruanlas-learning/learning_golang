package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go funcao1()
	go funcao2()

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}


func funcao1()  {
	fmt.Println("Função 1")
	wg.Done()
}

func funcao2()  {
	fmt.Println("Função 2")
	wg.Done()
}