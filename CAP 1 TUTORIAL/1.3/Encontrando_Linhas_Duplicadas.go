package main

import (
	//"fmt"
	"fmt"
	"os"

	//Responsavel pela leitura de dados
	//Bom para ler linha por linha. Evita ler tudo de uma vez e ideal para arquivos grandes
	"bufio"
)

func main() {
	//Counts = mapa
	//make cria e inicializa o mapa para uso
	//map[string]int -> chave - valor. Vai guardar a informação digitada e sua quantidade
	counts := make(map[string]int)

	//input vai guardar esse leitor de linhas
	//scanner é o leitor de linhas, vem da função NewScanner, que veio da importação do bufio
	//os.Stdio vai ser a fonte de dados(teclado) que meu scanner vai ler
	input := bufio.NewScanner(os.Stdin)

	//metodo de Scanner que vai premitir acessar os dados lidos de input
	//esse for vai verificar linha por linha de input, enquanto tiver linhas, ele vai retornar true
	//se ele retornar false, encerra o for, basicamente o while(sc.hasNextLine()) do java
	for input.Scan() {

		//a cada iteração do for, ele vai acessar um elemento do meu mapa counts, esse elemento seria a chave
		//a cada iteração, ele vai preencher aquela respectiva chave com o texto lido pelo input, da minha primeira linha
		//e o valor daquela chave, somado 1
		//quando ele procurar a chave, e achar uma com o mesmo nome, ele vai acessar aquela, ao inves de substituir ou criar uma nova, fazendo o incremento se for chave repetida
		counts[input.Text()]++
	}

	//esse for vai basicamente percorrer todos os elementos(chaves) do meu mapa counts, por meio do range
	//a cada chave, ele guarda o nome dela em lines e seu valor em n
	for line, n := range counts {

		//aqui ele verifica se o valor da minha chave atual é maior que 1, se for, ele executa o if
		if n > 1 {

			//print das palavras repetidas
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
