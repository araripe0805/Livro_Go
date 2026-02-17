package main

import (
	"fmt"
	"math"
)

// Quando se faz distanciaParaP := p.Distancia
// Voce não chama o metodo, que é perceptivel pela ausencia do parenteses
// Voce esta criando uma função nova já "presa" ao receptor p
// É como se o Go criasse algo como:
// func(q Point) float64 {return p.Distancia(q)}
// Ou seja, o receptor p fica guardado dentro da função
// Depois é só passar os outros argumentos
type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {
	p := Point{1, 2}
	q := Point{3, 4}

	distanciaParaP := p.Distance   //Valor de método
	fmt.Println(distanciaParaP(q)) //5

	//p.Distance virou uma função
	//Essa função já tem p armazenada
	//Só precisa de q agora
}
