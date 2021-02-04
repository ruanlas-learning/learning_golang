package main

import (
	"fmt"
)

type erroEspecial struct{
	msg string
}

func (err erroEspecial) Error() string {
	return fmt.Sprintf("Deu zica!!! %v", err.msg)
}

func main() {
	erro := erroEspecial{msg:"another error"}
	funcaoErroComoParametro(erro)
}

func funcaoErroComoParametro(err error) {
	fmt.Println(err)
}