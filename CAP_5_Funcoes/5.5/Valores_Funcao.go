package main

import (
	"fmt"
)

//Em go, funções podem ser atribuidas a variaveis, passadas como argumento, ser retornadas por outras funções, tem tipo proprio

// Função atribuida a variavel
// Não esta chamando a função, está guardando a função dentro da variavel f
// Posteriormente é chamada a função dentro a variavel f
func square(n int) int {
	return n * n
}

// O tipo de uma função
// O tipo da função é definido por sua assinatura: func(int) int
// Isso significa que ele recebe int e volta int
func negative(n int) int {
	return -n
}

// Erro no tipo da função
func product(m int, n int) int {
	return m * n
}

// Função como parametro
// parametros: f e seu tipo, n e seu tipo
// apply recebeu uma função
// ela quando executa, irá retornar chamando outra função, que por sua vez irá retornar seu valor para o primeiro retorno
func apply(f func(int) int, n int) int {
	return f(n)
}
func soma(n int) int {
	return n + n
}

func main() {
	f := square
	fmt.Println(f(3)) //9

	var a func(int) int
	a = negative
	fmt.Println(a(4)) //-4

	//Assinatura diferente == tipos diferentes
	var p func(int) int
	p := product

	resultado := apply(soma, 5)
	fmt.Println(resultado)
}
