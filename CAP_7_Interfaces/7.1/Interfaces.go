package main

import (
	"fmt"
)

// Tipos concretos definem como os dados são representados e quais operações ele tem
// Aqui é perceptivel como Pessoa é armazenada
type Pessoa struct {
	Nome string
}

// Interface = contrato
// Uma interface não diz o que é algo, ele diz o que aquele algo pode fazer
// Isso significa que qualquer um que tiver o metodo Falar() satisfaz a interface
type Falante interface {
	Falar() string
}

// Em Go, não é preciso dizer que tal tipo implementa uma interface, basta ter os métodos necessários e ele já estará implementado
// Tipo cachorro satisfaz a interface Falante graças ao método Falar
// Usando uma função para exibir o metodo Falar de qualquer integrante de Falante, isso permite deixar como parametro da função qualquer elemento que faça parte de Falante, não precisando fazer uma função para Falar de cada animal, dando flexibilidade e agrupando por comportamento, e não por estrutura
type Cachorro struct{}

func (c Cachorro) Falar() string {
	return "Au Au"
}

func fazerFalar(f Falante) {
	fmt.Println(f.Falar())
}

func main() {
	c := Cachorro{}
	fazerFalar(c)
}
