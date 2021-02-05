package main

import (
	"testing"
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
func TestSomaEmTabela2(t *testing.T) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},
		test{data: []int{10, 11, 12}, answer: 33},
		test{data: []int{-5, 0, 5, 10}, answer: 12},
	}

	for _, v := range tests {
		if res :=soma(v.data...); res != v.answer {
			t.Error("Expected:", v.answer, "Got:", res)
		}
	}
}