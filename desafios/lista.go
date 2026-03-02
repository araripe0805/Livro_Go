package main

import (
	"fmt"
)

type Lista interface {
	estaVazio() bool
	estaCheio() bool
	insereFinal(elemento int)
	insereInicio(elemento int)
	removeFinal()
	removeInicio()
	imprimirLista()
	inserePosicao(elemento, pos int)
	removePosicao(pos int)
	copiarLista() []int
	contemLista(elemento int) bool
	acesseLista(pos int) int
	alterarLista(pos, elemento int)
	indiceLista(elemento int) int
	subLista(a, b int) []int
	ordenarLista()
}

func main() {
	vetor := make([]int, 0, 5)
	ml := minhaLista{
		capacidade: 5,
		nElementos: len(vetor),
		vetor:      vetor,
	}

	ml.insereFinal(50)
	ml.insereFinal(20)
	ml.insereFinal(60)
	ml.insereFinal(10)
	ml.insereFinal(40)
	ml.imprimirLista()
	ml.ordenarLista()
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
	if ml.estaVazio() {
		fmt.Println("A lista está vazia")
		return
	}
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
		ml.insereInicio(elemento)
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

func (ml minhaLista) copiarLista() []int {
	if ml.estaVazio() {
		fmt.Println("A lista está vazia")
	}
	vetorCopia := make([]int, ml.nElementos, ml.capacidade)
	for i := 0; i < ml.nElementos; i++ {
		vetorCopia[i] = ml.vetor[i]
	}
	return vetorCopia
}

func (ml minhaLista) contemLista(elemento int) bool {
	if ml.estaVazio() {
		fmt.Println("A lista está vazia")
	}
	for i := 0; i < ml.nElementos; i++ {
		if ml.vetor[i] == elemento {
			return true
		}
	}
	return false
}

func (ml minhaLista) acesseLista(pos int) int {
	if ml.estaVazio() {
		fmt.Println("A lista está vazia")
	}
	if pos >= ml.nElementos {
		fmt.Println("A posição que você deseja acessar não existe")
	}
	elemento := 0
	for i := 0; i < ml.nElementos; i++ {
		if i == pos {
			elemento = ml.vetor[i]
		}
	}
	return elemento
}

func (ml *minhaLista) alterarLista(pos, elemento int) {
	if ml.estaVazio() {
		fmt.Println("A lista está vazia")
	}
	if pos >= ml.nElementos {
		fmt.Println("A posição que você deseja acessar não existe")
	}
	ml.vetor[pos] = elemento
}

func (ml minhaLista) indiceLista(elemento int) int {
	if ml.estaVazio() {
		fmt.Println("A lista está vazia")
	}
	for i := 0; i < ml.nElementos; i++ {
		if ml.vetor[i] == elemento {
			return i
		}
	}
	return -1
}

func (ml minhaLista) subLista(a, b int) []int {
	if ml.estaVazio() {
		fmt.Println("A lista está vazia")
	}
	if a < 0 {
		fmt.Println("Posição inicial invalida")
	}
	if b >= ml.nElementos {
		fmt.Println("Posição final invalida")
	}
	if b-a <= 0 {
		fmt.Println("Numeros de pos estão invalidos")
	}
	subLista := make([]int, 0, ml.capacidade)
	for i := a; i < b; i++ {
		subLista = append(subLista, ml.vetor[i])
	}
	return subLista
}

func (ml *minhaLista) ordenarLista() {
	if ml.estaVazio() {
		fmt.Println("A lista esta vazia")
	}
	for i := 0; i < ml.nElementos-1; i++ {
		for j := i + 1; j < ml.nElementos; j++ {
			if ml.vetor[i] > ml.vetor[j] {
				aux := ml.vetor[i]
				ml.vetor[i] = ml.vetor[j]
				ml.vetor[j] = aux
			}
		}
	}
}
