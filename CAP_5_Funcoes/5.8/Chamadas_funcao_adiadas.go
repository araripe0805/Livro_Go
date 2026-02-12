package main

import(
	"fmt"
)

//defer
//Adia a execução de uma função até o final da função atual
defer funcao()

//Internamente ele avalia os argumentos na hora
//Guarda a chamada numa pilha
//Ele executa só quando a função terminar
//Execução em ordem reversa(LIFO)
defer fmt.Println("fim")

//exemplo:
func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	//saidas
	//3,2,1(defer funciona como uma pilha(stack))
}

//Caso de fecha de recurso
//Duplicação de f.Close()
func lerArquivo() error {
	f, err := os.Open("Arquivo.txt")
	if err != nil {
		return err
	}

	//Se der erro aqui?
	if algumaCoisaErrada() {
		f.Close()
		return fmt.Errorf("erro")
	}

	f.Close()
	return nil
}

//Forma correta
//Forma mais limpa 
//Sem duplicação
//Sempre fecha o arquivo
func lerArquivo2() error {
	f, err := os.Open("Arquivo.txt")
	if err != nil {
		return err
	}

	//Colocado logo após abrir
	defer f.Close()

	if algumaCoisaErrada() {
		return fmt.Errorf("erro")
	}

	return nil
}