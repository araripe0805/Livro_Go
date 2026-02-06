package main

import (
	"fmt"
)

func main() {
	//String é um tipo de dado que contêm texto legível aos seres humanos
	//Normalmente representam texto
	s := "Hello"

	//não da para mudar um caractere diretamente
	s[0] = "h"

	//len(s) ler o numero de bytes, não de caracteres
	//s[i] acessa o  byte da posição desejada
	fmt.Println(len(s))
	fmt.Println(s[0])

	//Substring
	s = "hello, world"
	fmt.Println(s[0:5]) // "hello"
	fmt.Println(s[:5])  // "hello"
	fmt.Println(s[7:])  // "world"
	fmt.Println(s[:])   //"hello, world"

	//Concatenação
	a := "hello"
	b := "world"
	fmt.Println(a + "" + b)

	//Comparação
	fmt.Println("abc" < "abd") //true

	//Imutabilidade
	s = "left foot"
	t := s
	s += ", right foot"
	fmt.Println(s)
	fmt.Println(t)
}
