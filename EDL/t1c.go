package main

import (
	"fmt"
	"strings"
)

func contandoVogais(palavra string) int {
	contador := 0
	vogais := "aeiou"

	palavra = strings.ToLower(palavra) 

	for i := 0; i < len(palavra); i++ {
		letra := palavra[i]
		if strings.ContainsRune(vogais, rune(letra)) {
			contador++
		}
	}

	return contador
}

func main() {
	palavra := "Python"
	totalVogais := contandoVogais(palavra)

	fmt.Printf("Em'%s' tem número de vogais: %d .\n", palavra, totalVogais)
}
