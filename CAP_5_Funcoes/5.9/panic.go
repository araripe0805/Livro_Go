package main

import(
	"fmt"
)

//panic é um mecanismo que interrompe a execução normal
//Executa todos os defer da goroutine
//Mostra mensagem de erro
//Mostra a stack trace
//Encerra o programa

//tipos de ocorrencia de panic
//divisao por zero
//acesso fora do limite de um array
//desrefenrenciar ponteiro nil
//enviar para o canal fechado
//etc

//exemplo
func main() {
	x := 10
	y := 0
	fmt.Println(x/y)//panic
}

//exemplo
func array() {
	s := []int{3,2,4}
	fmt.Println(s[10])//panic
}

//O que acontece quando ocorre panic
//Execução normal para
//Go começa a "desenrolar" a stack
//Executa todos os defer
//Mostra o stack trace
//Encerra o programa
func f(x int) {
	fmt.Println("f(%d)\n",x+0/x) //panic quando x == 0
	defer fmt.Println("defer %d\n",x)
	f(x-1)
}
//f(3)
//f(2)
//f(1)
//defer 1
//defer 2
//defer 3
//panic: runtime error: integer divide by zero

//Os defer foram executados antes do programa morrer

//Voce ainda pode chamar o panic por conta propria
panic("algo deu muito errado")

//go não tem try catch. Logo, panic para bug grave, error para fluxo normal de erro