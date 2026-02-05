// Define o pacote
// Variaveis, funções... são validos para todo o pacote
package main

//Importações
import (
	"fmt"
)

func main() {
	//Variaveis
	//Opcao 1
	var idade int
	idade = 1

	//Opcao 2
	var idade2 int = 2

	//Go infere o tipo
	var idade3 = 3

	//forma curta, só funciona dentro de funções
	idade4 := 4

	fmt.Println(idade + idade2 + idade3 + idade4)

	//Constantes
	//Opcao 1
	const pi = 3.14

	//Opcao 2
	const pi2 float64 = 3.14

	fmt.Println(pi, pi2)

	//Type
	//Crio um novo tipo, para organizar melhor minhas variaveis
	type Idade int
	var x Idade = 20

	fmt.Printf("%T\n", x)

	//struct
	//Struct é um pacote de dados relacionados
	//criamos um tipo, chamado Pessoa, que terá algumas informações, como nome e idade, ou seja, todo x do tipo Pessoa tem Nome e Idade
	type Pessoa struct {
		Nome  string
		Idade int
	}

	//Criei uma variavel do tipo Pessoa, que terá seu nome e sua idade
	p := Pessoa{
		Nome:  "Inácio",
		Idade: 20,
	}

	fmt.Println(p)
	fmt.Println(p.Nome)
	fmt.Println(p.Idade)

	//t := new(Pessoa)

	//Mudando o valor de informações de uma variavel
	p.Nome = "Joao"
	fmt.Println(p.Nome)

	//funções
	//vai printar o retorno da função que está sendo chamada, passando como argumentos 2 e 3
	fmt.Println(soma(2, 3))

	fmt.Println(dividir(2, 2))
	fmt.Println(dividir(2, 0))
}

// Função que recebe 2 valores inteiros como parametro, e retornando um valor inteiro
func soma(x, y int) int {
	return x + y
}

// Função com 2 tipos de retorno
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divisão por zero")
	}
	return a / b, nil
}
