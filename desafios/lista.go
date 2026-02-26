package main

import (
	"fmt"
)

type Lista interface {
	estaVazio() bool
	estaCheio() bool
	insereFinal()
	insereInicio()
	removeFinal()
	removeInicio()
	imprimirLista()
	inserePosicao()
	removePosicao()
	copiarLista() []int
	contemLista() bool
	acesseLista() int
	alterarLista()
	indiceLista() int
	subLista() []int
	ordenarLista()
}

func main() {
	vetor := make([]int, 0, 5)
	ml := minhaLista{
		capacidade: 5,
		nElementos: len(vetor),
		vetor:      vetor,
	}

	ml.insereFinal(10)
	ml.insereFinal(20)
	ml.insereFinal(30)
	ml.insereFinal(40)
	ml.insereFinal(50)
	ml.imprimirLista()
	ml.removePosicao(2)
	ml.imprimirLista()
}

type minhaLista struct {
	capacidade int
	nElementos int
	vetor      []int
}

func (ml minhaLista) estaCheio() bool {
	return ml.capacidade == ml.nElementos
}

func (ml minhaLista) estaVazio() bool {
	return ml.nElementos == 0
}

func (ml minhaLista) imprimirLista() {
	fmt.Print("Lista: ")
	for i := 0; i < ml.nElementos; i++ {
		if i == ml.nElementos-1 {
			fmt.Println(ml.vetor[i])
			return
		}
		fmt.Printf("%d, ", ml.vetor[i])
	}
}

func (ml *minhaLista) insereInicio(elemento int) {
	if ml.estaCheio() {
		fmt.Println("A sua lista está cheia")
		return
	}
	ml.vetor = append(ml.vetor, 0)
	for i := ml.nElementos; i > 0; i-- {
		ml.vetor[i] = ml.vetor[i-1]
	}
	ml.vetor[0] = elemento
	ml.nElementos++
}

func (ml *minhaLista) removeInicio() {
	if ml.estaVazio() {
		fmt.Println("A lista está vazia")
		return
	}
	elementoRemovido := ml.vetor[0]
	for i := 1; i < ml.nElementos; i++ {
		ml.vetor[i-1] = ml.vetor[i]
	}
	ml.nElementos--
	fmt.Printf("Elemento removido: %d\n", elementoRemovido)
}

func (ml *minhaLista) insereFinal(elemento int) {
	if ml.estaCheio() {
		fmt.Println("A lista esta cheia")
		return
	}
	ml.vetor = append(ml.vetor, elemento)
	ml.nElementos++
}

func (ml *minhaLista) removeFinal() {
	if ml.estaVazio() {
		fmt.Println("A lista esta vazia")
		return
	}
	elementoRemovido := ml.vetor[ml.nElementos-1]
	fmt.Printf("Elemento removido: %d\n", elementoRemovido)
	ml.nElementos--
}

func (ml *minhaLista) inserePosicao(elemento, pos int) {
	if ml.estaVazio() {
		ml.insereFinal(elemento)
		return
	}
	if ml.estaCheio() {
		fmt.Println("A lista está cheia")
		return
	}
	for i := ml.nElementos - 1; i > pos; i-- {
		ml.vetor[i] = ml.vetor[i-1]
	}
	ml.vetor[pos] = elemento
	ml.nElementos++
}

func (ml *minhaLista) removePosicao(pos int) {
	if ml.estaVazio() {
		fmt.Println("A lista está vazia")
		return
	}
	elementoRemovido := ml.vetor[pos]
	for i := pos + 1; i < ml.nElementos; i++ {
		ml.vetor[i-1] = ml.vetor[i]
	}
	ml.nElementos--
	fmt.Printf("O elemento removido: %d\n", elementoRemovido)
}
