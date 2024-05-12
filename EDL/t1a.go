package main

import "fmt"

func separar(valores []int) (pares []int, impares []int) {
	for i := 0; i < len(valores); i++ {
		x := valores[i]

		if x%2 == 0 {
			pares = append(pares, x)
		} else {
			impares = append(impares, x)
		}
	}
	return pares, impares
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pares, impares := separar(numeros)

	fmt.Println("vetor Pares:", pares)
	fmt.Println("vetor Impares:", impares)
}
