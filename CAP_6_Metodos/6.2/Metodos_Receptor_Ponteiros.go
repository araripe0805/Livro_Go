package main

import(
	"fmt"
)

type Point struct {
	x, y float64
}

//p é uma copia, ele não possue o endereço do struct Point, é apenas uma copia, logo qualquer modificação em p, não altera os valores originais
func (p Point) Metodo() {
	p.x *= 4
	p.y *= 4
}

//Solução correta é usando ponteiros
func (p *Point) metodo2() {
	p.x *= 3
	p.y *= 2
}

func (p *Point) metodo3(x float64) {
	p.x *= x
	p.y *= x
}

//Como chamar metodos com ponteiros
p := Point{1,2}
pnt := &p
pnt.metodo3(2)

//ou
//p := Point{1,2}
//p.Metodo3(2)
