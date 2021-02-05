package main

import (
	"fmt"
	"exerciciosdocumentacao/cachorro"
)

func main() {
	idadeHumana := 1
	idadeCanina := cachorro.Idade(idadeHumana)
	fmt.Printf("A idade canina de um cachorro com %v ano é %v anos\n", idadeHumana, idadeCanina)
}