package main

import(
	"fmt"
)

//Recover é uma função que intercepta um panic e impede que o programa morra
	//Só funciona se for chamada dentro de uma função defer e ocorrer um panic em goroutine
	
	//Fluxo de um panic normal
	//função A
	//chama b
	//chama c
	//panic!
	//executa o defer c
	//executa o defer b
	//executa o defer a
	//programa termina

	//Fluxo com recover
	//função A
	//chama b
	//chama c
	//panic!
	//defer chama recover()
	//recover cancela o panic
	//função retorna normalmente
	//programa continua

	//um recover fora do defer, retorna nil
	defer func() {
		recover()
	}()

//exemplo
//programa morre
func a() {
	panic("erro fatal")
}

//exemplo
//Com recover
//panic foi chamado
//go guarda internamente o valor "algo deu errado"
//começa a executar os defer
//o recover() é chamado
//recover() olha: "há um panic ativo?"
//sim->devolve o valor guardado e guarda em r
//panic é cancelado
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado do panic:",r)
		}
	}()
	panic("algo deu errado")
	fmt.Println("isso nunca executa")
}