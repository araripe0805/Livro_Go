package main

import (
	"fmt"
)

func main() {
	//Inteiros podem ser com ou sem sinal
	//Com sinal(int), cobrem numeros negativos e positivos
	//Sem sinal(uint), cobre somente os numeros positivos
	//Compiladores normalmente interpretam esses 2 tipos com 32 ou 64 bits
	//Caso queira especificar a quantidade de bits pode-se usar: int8, int16, int32, int64, uint8...
	var a int = -10
	var b uint = 10

	fmt.Println(a, b)

	//O tipo rune é sinonimo para int32
	//Indica que esse valor é um codigo Unicode(Unicode code point)
	var c rune = 32
	var d int32 = 45

	//O tipo byte é sinonimo de uint8
	//Enfatiza que o valor é um dado bruto, e não uma quantidade numerica pequena
	var e byte = 80
	var f uint8 = 2

	//int é diferente de int8, int16...
	var g int32 = 1
	var h int64 = 2
	//Tipos diferentes não podem ser usados em operações aritmedicas
	soma := g + h

	//Forma de solucionar o problema de incompatibilidade dos tipos é fazendo a conversão
	//Normalmente a conversão de tipo não altera os valores
	//Mas a conversão que restrinja um inteiro maior em um valor menor, ou uma conversão de inteiro para ponto flutuante ou vice-versa, pode alterar o valor ou provocar perca de precição
	soma2 := int(g) + int(h)

}
