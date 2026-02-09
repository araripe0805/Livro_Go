package main

import (
	"fmt"
)

func main() {
	//Uma fatia(slice) é uma sequencia de tamanho variavel
	//Diferente de um Array, que já tem o tamanho fixo

	//Slice
	var a []int
	//Array
	var b [3]int

	//Uma fatia não guarda os dados diretamente
	//Fatia é só uma visão de um array
	//Internamente uma fatia possui isso
	type Slice struct {
		ptr *T  //Ponteiro para o array
		len int //tamanho
		cap int //capacidade (até onde pode crescer)
	}

	c := [6]int{10, 20, 30, 40, 50, 60}

	//Não é uma copia
	d := c[1:4]
	//O que acontece:
	//ptr -> &c[1] (20)
	//len = 3 (20,30,40)
	//cap = 5 (Do indice 1 até o fim do array)

	//Meu slice pode crescer enquanto houver espaço no array original
	//De 5 espaços disponiveis, ainda posso adicionar mais 2
	d = append(d, 99)
	//[20,30,40,99]

	//Caso não tenha mais espaço no array original
	//Cria um novo array, copia os valores e aponta a slice para o novo array
	//O array antigo continua existindo, mas a slice não aponta mais pra ele

	//Slice com função(sem uso de ponteiro)
	//Slice copia só a estrutura
	//o ptr aponta para o mesmo array
	//Então existe um ponteiro, mas ele está interno ao slice
	e := []int{1, 2, 3}
	altera(e)
	fmt.Println(e) //[99 2 3]

	//Bug, porque na função eu estou criando um novo array, e aí o novo slice só irá existir dentro da função
	//A forma de solucionar isso é usando um return
	f := []int{1, 2, 3}
	add(f)
	fmt.Println(f)

	//ptr = nil
	//len = 0
	//cap = 0
	var g []int

	//ptr != nil
	//len = 0
	//cap = 0
	h := []int{}

	//Operador de fatias
	//s[i:j] -> do indice i até j-1
	//s[:] -> tudo
	//s[:3] inicio até 2
	//s[2:] do indice 2 até o final
	months := [...]string{
		"", "January", "Febuary", "March", "April",
		"May", "June", "July", "August", "September", "October",
		"November", "December",
	}
	Q2 := months[4:7]      //[April may June]
	summers := months[6:9] //[June July August]

	//Slices compartilham memoria
	//Se mudar um slice, pode mudar um array e/ou outras slices
	//Isso acontece porque ambos apontam para um mesmo array
	i := []int{1, 2, 3, 4}
	j := i[1:4]
	j[0] = 99
	fmt.Println(i) //[1,99,3,4]

	//Criando slices
	//Slice literal
	k := []int{1, 2, 3, 4}

	//Com make
	l := make([]int, 3)    //len=3,cap=3
	m := make([]int, 3, 5) //len=3,cap=5

	//append com outro slice
	n := []int{1, 2}
	o := []int{3, 4}
	n = append(n, o)
	fmt.Println(n) //[1,2,3,4]

	//Copy
	p := []int{1, 2, 3}
	q := make([]int, len(p))
	copy(q, p) //[1,2,3]
}

func altera(e []int) {
	e[0] = 99
}

func add(f []int) {
	f = append(f, 99)
	//return f = append(f, 99)
}
