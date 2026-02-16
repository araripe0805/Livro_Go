package main

import (
	"fmt"
	"math"
)

//embedding é quando voce coloca uma struct dentro de outra sem dar o nome ao campo

//Exemplo sem embedding
type ColoredPoint struct {
	P Point
	Color string
}

//Com embedding
//Campo anonimo
type teste struct {
	Point
	Color string
}

//Exemplo
type Point struct {
	x, y float64
}

type ColoredPoint2 struct {
	Point
	Color string
}

//Acesso aos campos
var cp = Point{1,2}
//Sem embedding
//cp é a variavel que guarda o struct
//P está dentro de cp, e x está dentro de P
cp.P.x//1

//Com embedding
cp.x//1

//Exemplo mais completo
type valores struct {
	x,y float64
}

func (p *valores) multiplicacao(q float64) {
	p.x *= q
	p.y *= q
}

func (p valores) distancia (q valores) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

type cor struct {
	*valores
	Color string
}

func main() {
	p := cor{&valores{1,1}, "red"}
	q := cor{&valores{4,5}, "blue"}

	fmt.Println(p.distancia(*q.valores))

	p.multiplicacao(2)
}