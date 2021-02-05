package main

import (
	"fmt"
)

func main() {
	resultSoma := soma(1, 2, 3)
	resultMultiplica := multiplica(10, 10)
	fmt.Println(resultSoma)
	fmt.Println(resultMultiplica)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}