package main

import "fmt"

func buscaBinaria(nums []int, numero int) int {
	esquerda, direita := 0, len(nums)-1

	for esquerda <= direita {
		meio := (esquerda + direita) / 2

		if nums[meio] == numero {
			return meio
		} else if nums[meio] < numero {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}

	return -1
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := 7

	indice := buscaBinaria(numeros, x)

	if indice != -1 {
		fmt.Printf("O número %d está no índice %d\n", x, indice)
	} else {
		fmt.Printf("O número %d não está na lista\n", x)
	}
}
