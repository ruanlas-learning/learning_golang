package main

import (
	"fmt"
	"pacotes/pkg2"
)

// para executar rode em pkg1 => go run *.go
 
func main(){
	fmt.Println("Esta é uma funcao que está no pkg1 em main.go")
	Outro()
	pkg2.Pkg2()
}
