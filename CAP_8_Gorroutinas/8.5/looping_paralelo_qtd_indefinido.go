package main

//Supondo que você possui um canal que recebe numeros
//Mas você não sabe quantos numeros vão vim
//exemplo:
//numeros := make(chan int)
//vão ser enviados numeros para ela
//podem ser 3, 100, 1000
//Para cada numero recebido, serão criados goroutines que calculam o dobro
//e no final, somar todos os resultados

//O motivo para não usar len(slice):
//Os dados vem dessa forma:
//for n := range numeros {
//...
//}
//Só sabe que acabou a quantidade de numeros quando o canal é fechado

//A solução é usar sync.WaitGroup
//Ele serve para contar quantas goroutines estão rodando
//Esperar terminar

//Exemplo completo:
import (
	"fmt"
	"sync"
)

func main() {
	numeros := make(chan int)
	resultados := make(chan int)

	var wg sync.WaitGroup

	//Goroutine que consome os números
	go func() {
		//Incrementa contador(wg.Add(1))
		//Cria uma goroutine
		//Quando termina -> wg.Done()
		for n := range numeros {
			wg.Add(1)

			go func(num int) {
				defer wg.Done()
				resultados <- num * 2
			}(n)
		}

		//Quando o canal numeros é fechado, o for range termina
		//Esperamos todas as goroutines terminarem para fechar resultados
		wg.Wait()
		close(resultados)
	}()

	//Goroutine que envia números (quantidade desconhecida)
	go func() {
		for i := 1; i <= 5; i++ { // Imagine que isso poderia ser qualquer quantidade
			numeros <- i
		}
		close(numeros)
	}()

	//Somando resultados
	total := 0
	for r := range resultados {
		total += r
	}
	fmt.Println("Total", total)
	//Saida:
	//Total: 30
	//2 + 4 + 6 + 8
}

//Enviador ->
//Numeros ->
//Loop consumidor ->
//cria goroutines ->
//goroutines enviam em resultados ->
//wg.Wait() ->
//close(resultados) ->
//loop principal termina
