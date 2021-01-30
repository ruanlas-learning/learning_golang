package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var cont int64 = 0
	totalGoRoutines := 600

	callGoRoutines(totalGoRoutines, &cont)

	wg.Wait()

	fmt.Println()
	fmt.Println("Total de goroutines:", totalGoRoutines)
	fmt.Println("Valor do contador:", cont)
}

func callGoRoutines(qtdCalls int, cont *int64) {
	wg.Add(qtdCalls)
	for i := 0; i < qtdCalls; i++ {
		go func (cont *int64, id int)  {
			fmt.Println("GoRoutine nº: ", id, "\tValor do [cont]:", *cont)
			
			atomic.AddInt64(cont, 1)

			// temp := *cont
			runtime.Gosched()
			// temp++
			// *cont = temp

			wg.Done()
		}(cont, i)
	}
}