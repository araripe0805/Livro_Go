package main

import (
	"fmt"
	"os"
	"time"
)

// O select serve para esperar multiplas operações com canais ao mesmo tempo
// Ele é como um switch, mas em vez de comparar valores, ele espera comunicações em canais
func main() {
	//O select serve para esperar multiplas operações com canais ao mesmo tempo
	//Ele é como um switch, mas em vez de comparar valores, ele espera comunicações em canais
	//select {
	//case <-ch1:
	//recebeu de ch1
	//case x := <-ch2:
	//recebeu de ch2
	//case ch3 <- y:
	//enviou para ch3
	//default:
	//executa se nenhum canal estiver pronto
	//}

	//Ele espera:
	//Espera até algum canal estar pronto
	//Executa apenas um dos cases
	//Se vários estiverem prontos, escolhe aleatoriamente
	//Se nenhum estiver pronto, bloqueia(a menos que exista default)

	//Exemplo 1: Esperando dois eventos
	abort := make(chan struct{})

	//Goroutine que espera o usuario apertar Enter
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Espera 5 segundos ou Enter...")

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("tempo esgotado")
	case <-abort:
		fmt.Println("Abortado pelo usuario")
	}
	//Ou espera 5 segundos, ou o usuario aperta Enter
	//time.After cria um canal que envia um valor depois de 5seg
	//abort recebe valor se o usuario aperta Enter
	//O select espera o primeiro evento acabar

	//Por que não fazer isso:
	//<-time.After(5*time.Second)
	//<-abort
	//O primeiro <- bloqueia até terminar
	//O segundo nunca será alcançado se o primeiro bloquear

	//Exemplo 2 — Contador com cancelamento
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	fmt.Println("Contagem iniciada. Pressione Enter para abortar.")

	for i := 5; i > 0; i-- {
		fmt.Println(i)

		select {
		case <-ticker.C:
			//passou 1 segundo
		case <-abort:
			fmt.Println("Abortando")
			return
		}
	}
	//ticker.C envia um valor a cada 1 segundo
	//select decide: continuar contagem ou abortar

	//Exemplo 3 — Canal com buffer 1 (o exemplo do livro)
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
	//Canal tem 1 buffer
	//Se estiver  vazio -> só pode enviar
	//Se estiver cheio -> só pode receber
}
