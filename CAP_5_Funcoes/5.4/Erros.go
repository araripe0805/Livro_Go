package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"string"
	"time"
)

func main() {
	//Erros fazem parte do fluxo normal do programa
	//Ele não possui excessões como java ou python
	//Ele usa valores de retorno

	//Tipos de funções quanto a erro

	//Funções que não podem falhar
	string.Contains("golang", "go")
	strconv.FormatBool(true)
	//Essa função sempre funciona

	//Funções que só falham se houver bug
	time.Date(2026, 2, 11, 10, 0, 0, 0, nil)
	//Se o ultimo for um nil,ocorre um panic
	//Isso indica erro do programador

	//Funções que podem falhar normalmente:
	//Ler arquivos, fazer requisição HTTP, conectar banco, abrir conexão

	//Tipo error
	//se for nil, não possui erro, sucesso
	//qualquer valor diferente de nil, indica falha
	type error interface {
		Error() string
	}

	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	fmt.Println(resultado)
}

// Exemplo
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

//Estrategias de tratamento de erro

//1 estrategia-Propagar	o erro
//Se sua função não sabe resolver o erro, ele repassa para quem chamou
//Isso cria uma cadeia de responsabilidade
//ReadFile não tenta resolver o problema
//Apenas devolve o erro
func lerArquivo(nome string) ([]byte, error) {
	dados, err := os.ReadFile(nome)
	if err != nil {
		return nil, err//ou uma mensagem de erro
	}
	return dados,nil
}

//2 estrategia-Tentar novamente (Retry)
//Quando o erro pode ser temporario
//Quando o servidor está fora do ar, falha momentanea de rede, banco ocupado
func conectar() error {
	for i := 0; i < 3; i++ {
		err := tentarConectar()
		if err == nil {
			return nil
		}
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("Não foi possivel conectar após 3 tentativas")
}

//3 estrategia-Encerrar o programa
//Somente no main ou camadas mais altas
//Bibliotecas não podem ser encerradas
//Isso mostra o erro e encerra o programa
func main() {
	err := iniciarServidor()
	if err != nil {
		log.Fatalf("Erro fatal: %v", err)
	}
}

//4 estrategia-Registrar erros e continuar
//Quando o erro não é critico
//Quando falhou em enviar alguma estatistica, falhou cache
//Programa continua rodando
if err := enviarMensagem(); if err != nil {
	//Mesma função do fmt, mas adiciona a data e a hora
	log.Printf("Falha ao enviar a mensagem: %v", err)
}

//5 estrategia-Ignorar erro
//Deve ser raro e consciente
os.RemoveAll(dir)//erro ignorado propositalmente
//O sistema limpa diretorios temporarios depois

//Padrão de estrutura em Go
func exemplo() error {
	//Checa erro
	if err := passo1(); err != nil {
		return err
	}

	//Checa o erro
	if err := passo2(); err != nil {
		return err
	}

	//Logica principal
	fmt.Println("Tudo certo")
	return nil
}
	
