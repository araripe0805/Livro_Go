package main

import(
	"fmt"
)

//Estrutura básica de uma função
func nome(parametros) (resultados ou retornos) {
	conteudo
}

//Exemplo
func soma(a, b int) int {
	return a+b
}

//Aplicação
fmt.Println(soma(3,2))

//Argumentos são usadas ao chamar a função
//Parametro são variaveis locais, recebidas na função e irão ser usadas na respectiva

//Se retornar mais de um valor
func nomeIdade(a int,b string) (int,string) {
	return a, b
}

fmt.Println(nomeIdade(20, "inacio"))

//retornar nada
func imprimir() {
	fmt.Println("Imprimindo mensagem...")
}

imprimir()

//Passagem de valor
//Go sempre copia
func x (y int) {
	y = 100
}

a := 10
x(a)
fmt.Println(a)//10

//O valor de a não é alterado, apenas a variavel local foi modificada
//Diferente de tipo um slice sendo usado em uma função, ele pode ser alterado, já que aponta para um array
