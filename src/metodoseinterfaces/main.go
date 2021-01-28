package main

import (
	"fmt"
)

type pessoa struct{
	nome string
}

// func (p pessoa) falar() {
// 	fmt.Println("Olá, meu nome é",  p.nome)
// }

func (p *pessoa) falar() {
	fmt.Println("Olá, meu nome é",  p.nome)
}

type humanos interface{
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	p1 := pessoa{nome: "Paulo"}

	dizerAlgumaCoisa(&p1)

	// dizerAlgumaCoisa(p1)


}