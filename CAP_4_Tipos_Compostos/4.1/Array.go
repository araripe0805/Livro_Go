package main

import (
	"fmt"
)

func main() {
	//Array é uma sequencia de tamanho fixo
	//Todos os elementos possuem o mesmo tipo
	//Em go, Arrays não são tão utilizados, sendo substituido mais vezes por slices

	//Declaração do meu array
	//O numero representa o tamanho do array
	var a [3]int

	//Exibi o primeiro elemento
	fmt.Println(a[0])

	//Exibi o ultimo elemento
	fmt.Println(a[len(a)-1])

	//Percorrendo todo o Array
	for i, v := range a {
		fmt.Println(v, i)
	}

	//Inicializando com numeros já definidos
	q := [3]int{1, 2, 3}

	//Tamanho inferido com ...
	r := [...]int{4, 5, 6}

	//Tamanho faz parte do tipo
	//c := [2]int{1,2}
	//d := [1]int{4}
	//c != d

	//Inicialização com Indices
	//[10,0,0,40,0]
	//Os indices não precisam estar em ordem
	//Indices que não forem informados recebem o valor 0
	c := [5]int{
		0: 10,
		3: 40,
	}

	//Inicializando com indices e inferencia de tamanho
	//[0,0,2]
	d := [...]int{
		2: 7,
	}

	//Misturando valores sequenciais e de indices
	//[1,2,4,0,0,20,8]
	e := [...]int{
		1, //Sequencia
		2, //Sequencia
		4, //Sequencia
		5: 20, //Indice
		8, //Sequencia
	}

	//Exemplo enum + array
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	//Esses indices funcionam, já que essas constantes são inteiros
	symbols := [...]string{
		USD: "$",
		EUR: "€",
		GBP: "£",
		RMB: "¥",
	}

	fmt.Println(symbols[RMB])

	//A função não altera os valores do array, a não ser que eu retorne o valor dela
	var f [3]int
	fmt.Println(array(f))

	//Ponteiro irá apontar para o endereço do array, por isso que ele consegue usar/alterar os valores mexidos na função, fora dela, sem usar um return
	var g [4]int
	alterar(&g)

	//Inicia com 4 valores 0
	//[0 0 0 0]
	//Alterando um indice depois
	//[9 0 10 0]
	var h [3]int
	h[2] = 10
	h[0] = 9
}

func array(f [3]int) [3]int {
	f[0] = 99
	return f
}

func alterar(g *[4]int) {
	g[0] = 4
}
