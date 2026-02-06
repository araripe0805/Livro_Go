package main

import (
	"fmt"
)

func main() {
	//Tipos Nomeados
	//Essa declaração cria um novo tipo , baseado em outro já existente(subjacente)
	//Serve para dar um significado melhor para o tipo das suas declarações
	//Não é o mesmo que um int, ele é do tipo Idade, que foi baseado em um int
	type Idade int

	//x e y possuem ambos inteiros como valor, mas são de tipos dfiferentes
	//Tipos Nomeados não são automaticamente compativeis
	type ano int
	var x Idade = 20
	var y ano = 2026

	fmt.Println(x, y)

	//tipo y != tipo x; Logo x = y não é possivel
	//x = y

	//Forma de converter o tipo, para ser compativel para outras operações
	x = Idade(y)

	fmt.Println(x, y)

	//Tipos Nomeados com struct
	//Agora Pessoa é um tipo proprio(tipo nomeado), assim como Idade e ano
	//Struct é um tipo composto, que pode ser o tipo subjacente de um tipo nomeado
	//Ex: Idade(tipo nomeado) -> int(tipo subjacente) ; Pessoa(tipo nomeado) -> struct{Nome string Peso int}(tipo subjacente)
	type Pessoa struct {
		Nome string
		Peso int
	}
}
