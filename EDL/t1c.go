package main

import (
	"fmt"
	"strings"
)

func contarVogais(palavra string) int {
	contador := 0
	vogais := "aeiou"

	palavra = strings.ToLower(palavra)

	for i := 0; i < len(palavra); i++ {
		letra := palavra[i]
		if strings.ContainsAny(vogais, string(letra)) {
			contador++
		}
	}

	return contador
}

func main() {
	palavra := "LinguAGEM go"
	totalVogais := contarVogais(palavra)

	fmt.Printf("Na palavra '%s' existem %d vogais.\n", palavra, totalVogais)
}
