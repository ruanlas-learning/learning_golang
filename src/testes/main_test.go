package main

import (
	"testing"
	"fmt"
)

type test struct{
	data []int
	answer int
}

func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

// // esse falha
// func TestSoma2(t *testing.T) {
// 	teste := soma(3, 2, 1)
// 	resultado := 2
// 	if teste != resultado {
// 		t.Error("Expected:", resultado, "Got:", teste)
// 	}
// }

func TestMultiplica(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

// // esse falha
// func TestMultiplica2(t *testing.T) {
// 	teste := multiplica(10, 10)
// 	resultado := 101
// 	if teste != resultado {
// 		t.Error("Expected:", resultado, "Got:", teste)
// 	}
// }


func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},
		test{data: []int{10, 11, 12}, answer: 33},
		test{data: []int{-5, 0, 5, 10}, answer: 10},
	}

	for _, v := range tests {
		if res :=soma(v.data...); res != v.answer {
			t.Error("Expected:", v.answer, "Got:", res)
		}
	}
}

// esse falha
// func TestSomaEmTabela2(t *testing.T) {
// 	tests := []test{
// 		test{data: []int{1, 2, 3}, answer: 6},
// 		test{data: []int{10, 11, 12}, answer: 33},
// 		test{data: []int{-5, 0, 5, 10}, answer: 12},
// 	}

// 	for _, v := range tests {
// 		if res :=soma(v.data...); res != v.answer {
// 			t.Error("Expected:", v.answer, "Got:", res)
// 		}
// 	}
// }

// teste como exemplo
func ExampleSoma() {
	fmt.Println(soma(3, 2, 1))
	// Output: 6
}

func ExampleMultiplica() {
	fmt.Println(multiplica(3, 2, 2))
	fmt.Println(multiplica(1, 2, 2))
	fmt.Println(multiplica(2, 2, 2))
	// Output: 
	// 12
	// 4
	// 8
}


/* 
	go test testes/ -bench . 
		* Executa todos os benchmarks
	go test testes/ -bench Soma
		* Executa o benchmark apenas da função soma
*/

func BenchmarkSoma(b *testing.B) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},
		test{data: []int{10, 11, 12}, answer: 33},
		test{data: []int{-5, 0, 5, 10}, answer: 10},
	}

	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			soma(v.data...)
		}
	}
}


func BenchmarkMultiplica(b *testing.B) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},
		test{data: []int{10, 11, 12}, answer: 33},
		test{data: []int{-5, 0, 5, 10}, answer: 10},
	}

	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			multiplica(v.data...)
		}
	}
}