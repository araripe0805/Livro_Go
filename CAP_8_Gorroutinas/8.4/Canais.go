package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//Goroutines são atividades concorrentes
	//channels são o meio de comunicação entre elas
	//Um canal transporta valores de um tipo especifico
	var ch chan int     //canal de int
	ch = make(chan int) //cria o canal

	//Enviar e receber usa o operador <-
	//ch <- 10  //Envia
	//x := <-ch //Recebe
	//<-ch //recebe e descarta

	//Canais são referencias, como mapas e slices

	//Canais sem buffer(Unbuffered)
	//ch := make(chan int)
	//Um envio bloqueia até existir um receptor
	//Uma recepção bloqueia até existir um envio
	go func() {
		ch <- 42 //Bloqueia até alguém receber
	}()

	valor := <-ch
	fmt.Println(valor)
	//Goroutine tenta enviar 42
	//Ela bloqueia
	//Main recebe
	//Ambas continuam
	//Isso cria sincronização garantida (happens-before)

	//Canais como sinalização de evento
	//As vezes não é necessario enviar dados-- Só sinalizar
	//Usa-se struct{} porque ocupada 0 bytes
	//O valor passado no canal não importa, só importa o evento
	done := make(chan struct{})
	
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine terminou")
		done <- struct{}{} 
	}()

	fmt.Println("Main esperando...")
	<-done
	fmt.Println("Main continua aqui")
	//1.Main esperando
	//2. 2seg
	//3. Goroutine terminou
	//4. Main continua
	//Main não passa de <-done antes do envio acontecer

	//Fechando canais
	//close(ch)
	//Enviar(Falha) -> panic
	//Receber(sucesso) -> continua funcionando
	//Após esvaziar -> retorna valor zero

	//Usando range com canais
	//Ele vai receber valores do canal até ele ser fechado
	//Ou seja, o loop só acaba quando o canal é fechado
	for v := range ch {
		fmt.Println(v)
	}

	//Pipeline = saida de uma goroutine -> entrada de outra
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for i := 0; i < 5; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	//Squarer
	go func() {
		for n := range naturals {
			squares <- n * n
		}
		close(squares)
	}()

	//Printer
	for s := range squares {
		fmt.Println(s)
	}
	//counter -> squarer -> printer
	//0
	//1
	//4
	//9
	//16

	//Canais unidirecionais
	//Erro apenas para ficar melhor na explicação
	func produtor(out char<- int) {
		out <- 10
	}
	//chan<- apenas envia

	func consumidora(in <-char int) {
		fmt.Println(<-in)
	}
	//<-chan só recebe

	//Canais com Buffer
	ch = make(chan strings, 2)
	ch <- "A"
	ch <- "B"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//Os envios não bloqueiam até encher o buffer
	//Só bloqueia quando ultrapassa a capacidade
}
