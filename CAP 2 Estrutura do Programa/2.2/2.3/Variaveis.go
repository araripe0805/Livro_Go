package main

import (
	"fmt"
)

func main() {
	//Declaração sem valor inicial
	//int = 0; float64 = 0.0; string = ""; bool = false; ponteiros = nil
	var x int

	//Declaração com valor inicial
	var y int = 10

	//Com inferencia de tipo
	var z = 11

	fmt.Println(x, y, z)

	//Declarações de Variaveis curtas
	//Curta, só funciona em funções, Go infere o tipo
	a := 10
	b := "Go"
	c := true

	fmt.Println(a, b, c)

	//Declaração multipla
	d, e := 1, 2

	fmt.Println(d, e)

	//Shadowing(sombreamento de variaveis)
	//Variavel interna declarada não altera a externa declarada
	//Fora de condicionais ou de laços de repetição daria erro, pois não se pode declarar 2 variaveis com o mesmo nome e mesmo nivel
	//= != := (são diferentes, = recebe e := declara e recebe )
	f := 10
	if true {
		f := 20
		fmt.Println(f)
	}
	fmt.Println(f)

	//Ponteiros
	//Ponteiros guardam o endereço de uma variavel
	//g é uma variavel, e h irá guardar o endereço de h (&h), sendo um ponteiro, apontando para g
	g := 30
	h := &g

	//Valor apontado(*h)
	fmt.Println(*h)

	//Alterar o valor do ponteiro, altera o valor da variavel original
	//h = 20 não compilaria ,  h é um ponteiro, e não uma variavel, que irá receber um novo valor
	*h = 20
	fmt.Println(g)

	//New
	//new cria um ponteiro para o valor zero de algum tipo
	//Equivalente a: var x int; p := &x
	p := new(int)
	*p = 5

	fmt.Println(*p)

	//Tipo Pessoa
	type Pessoa struct {
		Nome  string
		Idade int
	}

	//Criei uma Pessoa, mas todos os valores dos campos estão como zero
	i := new(Pessoa)

	fmt.Println(i.Idade, i.Nome)

	//Por que eu consigo usar i.informação? Pode-se usar tanto a forma abaixo, quanto (*i).Nome, são equivalentes
	i.Nome = "Inacio"
	i.Idade = 20

	fmt.Println(i.Idade, i.Nome)

	//A diferença de usar Struct com new e sem o new, é basicamente que usando o new, se trata de um ponteiro, e não uma variavel

	//Outra forma de criar um Struct como ponteiro, sem o new
	//Mais usado quando você não quer iniciar o seu ponteiro com valores zerados
	j := &Pessoa{
		Nome:  "João",
		Idade: 19,
	}

	fmt.Println(j.Nome, j.Idade)

	//Tempo de Vida das Variaveis
	//Tempo de vida depende do escopo

	//Variaveis de nivel de pacote
	//Existem durante toda a execução do programa
	//Acessivel por todo o pacote
	//var k int

	//Variaveis Locais
	//Vive apenas dentro de uma função
	//São destruidas se sair do escopo
	//func soma() {
	//	x := 10
	//}

	//Ponteiros e tempo de vida
	//escape Analysis
	//func cria() *int {
	//	x := 10
	//	return &x
	//}

}
