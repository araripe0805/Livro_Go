package main

import "fmt"

func main() {
	//Constantes são valores conhecidos em tempo de compilação
	//Nunca mudam durante a execução
	//Podem ser booleanos, string e numeros
	//Não podem ser alterados depois que declarados
	const pi = 3.14159

	//Vantagens
	//Avaliadas antes do programa rodar
	//Mais rapidas

	//Grupo de constantes
	const (
		a = 1
		b = 2
		c = 3
	)

	//Se omitir o valor, o Go repete o valor anterior
	const (
		d = 1
		e
		f = 2
		g
	)
	fmt.Println(d, e, f, g) //1,1,2,2

	//iota-gerador de constantes
	//iota é um contador automatico que só existe dentro de um bloco const
	//Ele sempre começa com 0 e vai incrementando
	type weekday int
	const (
		Sunday    weekday = iota //0
		Monday                   //1
		Tuesday                  //2
		Wednesday                //3
		Thursday                 //4
		Friday                   //5
		Saturday                 //6
	)

	//Muito usado quando zero não deve ser um valor valido
	const (
		_ = iota // ignora 0
		A        //1
		B        //2
	)

	fmt.Println(A, B)
}
