package main

import (
	"fmt"
)

// Declarações fora de funções são visiveis em todo o pacote
var x = 10

// Variaveis declaradas dentro de uma função só existem dentro daquela propria, a não ser que tenha um retorno, ou chame outra função, passando a variavel como parametro
func main() {
	fmt.Println(x)
	//y := 29 -> erro
}

func outro() {
	//fmt.Println(y) -> erro
}

func media() {
	media := 8
	if media >= 6 {
		aprovacao := "aprovado"
		fmt.Println(aprovacao)

	} else {
		aprovacao := "reprovado"
		fmt.Println(aprovacao)
	}

	//Cada {} cria um novo escopo, logo aprovacao não pode ser usado fora do if e else, porque não foi declarafo fora desse escopo(if e else)
	//fmt.Println(aprovacao) -> erro

	//Escopo no for funciona da mesma forma que o if e else
	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}
	//fmt.Println(i)

	//Escopo no switch case funciona da mesma forma também
	//switch x := 10; x {
	//	case 10:
	//		fmt.Println(x)
	//}
	//fmt.Println(x)
}
