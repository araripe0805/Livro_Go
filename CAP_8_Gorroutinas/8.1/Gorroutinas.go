package main

import (
	"fmt"
	"time"
)

//Programação concorrente significa executar multiplas atividades ao mesmo tempo dentro do mesmo programa
//Varias atividades executam concorrentemente
//Elas não compartilham variaveis diretamente
//elas se comunicam trocando mensagens por canais
//"Não compartilhar memoria para se comunicar, se comunicar para compartilhar memoria"

// Gorrotina é uma função que executa concorrentemente com outras funções
func main() {
	//funcao() //chamada normal(bloqueia até terminar)
	//go funcao() //executa funcao() em paralelo(não espera terminar)

	//Quando o programa inicia , apenas a função main() está rodando, chamada de gorrotina principal
	//Quando main()termina: Todas as outras gorrotinas são encerradas  automaticamente
	//exemplo:
	//go tarefa() cria uma nova gorrotina e o main continua executando
	go tarefa()
	fmt.Println("Main terminou")
	//Se não for usado Sleep, o programa termina antes da tarefa acabar
	time.Sleep(6 * time.Second)

	//Uma gorrotina não pode "matar" outra, e deve-se ter cuidado ao usar variaveis compartilhadas

	//exemplo 2: Duas tarefas concorrentes
	go contador("a")
	go contador("b")
	time.Sleep(3 * time.Second)

	//Saida:
	//A 1
	//B 1
	//A 2
	//B 2
	//...
	//Ordem pode variar a cada execução
}

func funcao()

func tarefa() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Executando tarefa:", i)
		time.Sleep(time.Second)
	}
}

func contador(x string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(x, i)
		time.Sleep(500 * time.Millisecond)
	}
}
