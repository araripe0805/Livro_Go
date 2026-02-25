package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numeroDigitado int
	numeroSorteado := rand.Intn(101)
	tentativas := 1

	for {
		fmt.Print("Digite um número: ")
		fmt.Scan(&numeroDigitado)

		opcao := numeroSorteado - numeroDigitado
		switch {
		case opcao == 0:
			fmt.Printf("Parabens!!! Você acertou o numero %d em %d tentativas\n", numeroSorteado, tentativas)
			return
		case opcao >= 1:
			fmt.Printf("O número digitado é menor que o numero secreto\nTentativas: %d\n", tentativas)
		case opcao <= -1:
			fmt.Printf("O numero digitado é maior que o numero secreto\nTentativas: %d\n", tentativas)
		}
		tentativas++
	}

}
