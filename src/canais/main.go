package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go func(){
		canal <- 42
	}()

	fmt.Println(<-canal)
}


/* Não funciona!!!! O canal recebe uma atribuição e fica   
	esperando outro pegar o valor */
// func main() {
// 	canal := make(chan int)

// 	canal <- 42

// 	fmt.Println(<-canal)
// }