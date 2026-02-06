// Um pacote agrupa codigos relacionados, é a unidade básica de organização
// Sempre deve ser a primeira linha de codigos em Go
// Codigos no mesmo diretorio, e com o mesmo nome de pacote, fazem parte do mesmo pacote
// Posso ter 2 pacotes diferentes com o mesmo nome, desde que estejam em diretorios diferentes
// Funções e variaveis do mesmo pacote, podem ser compartilhadas
package main

//Pacotes são importados pelo caminho de importações
//Você pode importar os seus proprios pacotes, não necessariamente os já feitos pelo Go
import (
	"fmt"
	"math"
)

func main() {
	//Funções, metodos, variaveis... Que você quiser acessar de pacotes importados, devem ter essa estrutura, prefixo(nome do pacote) + sufixo(função,... usada)
	fmt.Println("Hello World!")
	fmt.Println(math.Sqrt(4))
}

// Funções com a primeira letra maiuscula, são acessiveis fora do pacote
func Publica() {

}

// Funções com a primeira letra minuscula, não são acessiveis fora do pacote
func privada() {

}

//Ordem de inicialização: Variaveis do pacote -> funções init -> main()
//Função init é uma função executada automaticamente, pode existir em qualquer arquivo do pacote, não recebe parametros e não retorna valor
