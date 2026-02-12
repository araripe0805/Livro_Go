package main

import(
	"fmt"
)

//Função anonima é uma função sem nome
//Pode ser atribuida a variaveis
//pode ser passada como argumento
//pode ser retornada por outra função
func(parametros) return {
	corpo
}

//exemplo
//Foi criada uma função anonima
//Essa função foi guardada em uma variavel f
//Chamada normalmente
func main() {
	f := func(a int, b int) int {
		return a + b
	}
	fmt.Println(f(2,3))
}

//Usando a função como parametro
//Aqui não foi criada variavel para guardar a função
//Ela é passada direta como argumento, comum quando a função só será usada uma vez e não será reutilizada
...strings.Map(func(r rune) rune {
	return r + 1
}, "HAL-9000")

//Closures
//É uma função que captura variaveis do escopo externo
//Mantem essas variaveis vivas mesmo após o retorno da função
//Normalmente as variaveis de uma função são criadas como stack, pois normalmente possuem um uso apenas local
//Mas quando você retorna uma função e essa função captura uma variavel externa, ela passa para a heap, assim mantendo sempre os valores atualizados da variavel
//Lembrando que as variaveis normalmente guardam os retornos, logo, f irá guarda justamente a função que retorna x
func squares() func int {
	var x int

	return func() int {
		x++
		return x*x
	}
}

//Irá guardar a função anonima 
f := squares()
fmt.Println(f())
fmt.Println(f())
fmt.Println(f())
fmt.Println(f())

//topoSort(busca de profundidade)
//Por que não fazer assim: visitAll := func...? 
//Porque quando vai ocorrer a recursividade a variavel tem que estar definida no momento da criação
//No exemplo abaixo, primeiro a variavel é declarada e depois atribuida
var visitAll func(items []string)

visitAll = func(items []string) {
	for _, item := range items {
		if !seen[item] {
			seen[item] = true
			visitAll(m[item])
			order = append(order, item)
		}
	}
}

//Extract(captura variavel)
//A função anonima usa a variavel links
//Ela não recebeu links como parametro, ela capturou o escopo externo
var links []string

visitNode := func(n *html.Node) {
	links = append(links, link.String())
}