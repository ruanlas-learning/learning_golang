package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	cont := 0
	totalGoRoutines := 600

	callGoRoutines(totalGoRoutines, &cont)

	wg.Wait()

	fmt.Println()
	fmt.Println("Total de goroutines:", totalGoRoutines)
	fmt.Println("Valor do contador:", cont)
}

func callGoRoutines(qtdCalls int, cont *int) {
	wg.Add(qtdCalls)
	for i := 0; i < qtdCalls; i++ {
		go func (cont *int, id int)  {
			fmt.Println("GoRoutine nº: ", id, "\tValor do [cont]:", *cont)
			temp := *cont
			runtime.Gosched()
			temp++
			*cont = temp

			wg.Done()
		}(cont, i)
	}
}