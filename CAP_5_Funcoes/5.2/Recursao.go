package main

import (
	"fmt"
)

//Recursão é quando uma função chama ela mesma
//É necessario um codigo base(quando parar) e um passo recursivo(chamada)
//Sem esses 2 pontos, a função ou se torna um loop infinito, ou ele não irá retornar o valor desejado

//Exemplo de recursividade

//Fatorial
//Definição recursiva: n! = n x (n-1)
//codigo base: 1! = 1

func fatorial(x int) int {
	//Codigo base
	if x == 1 {
		return 1
	}
	//Chamada recursiva
	return x * fatorial(x-1)
}

func main() {
	fmt.Println(fatorial(5))
}

//Fluxo mental:
//5 * fatorial(4)
//4 * fatorial(3)
//3 * fatorial(2)
//2 * fatorial(1)
//1 -> codigo base
//2 * 1
//3 * 2
//4 * 6
//5 * 24
//= 120

//Recursão é per
