package main

import (
	"fmt"
	"sort"
)

func main() {
	//Mapa é uma estrutura chave -> valor
	//"alice" -> 31
	//"bob" -> 32
	//Em go: map[k]v
	//exemplos: map[string]int

	//Map é uma hash table
	//Acesso rápido
	//Não tem ordem
	//Você busca não pelo indice, busca pelo nome da chave
	ages := make(map[string]int)

	ages["Alice"] = 32
	ages["Bob"] = 31

	//Criando mapas

	//Com make
	peso := make(map[string]int)

	//literal
	altura := map[string]int{
		"Alice": 32,
		"Bob":   30,
	}

	//Ambos são equivalentes

	//Acessando e alterando valor
	//Se buscar uma chave que não existi, o valor é zero
	ages["Alice"] = 20

	//Incrementando valores que podem não existir
	//0 + 1
	ages["Carlos"]++

	//Percorrendo um mapa(map)
	//Ordem não garantida, já que não existi indice
	for name, age := range ages {
		fmt.Println(name, age)
	}

	//Iterar em ordem(Ordenando chaves)
	//Cria um slice
	names := make([]string, 0, len(ages))

	//Vai percorrer todas as chaves de ages, com o range
	//essas chaves de ages irão ser guardadas em name
	//e age vai guardar o respectivo valor
	for name, age := range ages {

		//e no slice names ele irá guardar as chaves(name)
		names = append(names, name)
	}

	//Irá ordenar os slices pela ordem alfabetica
	sort.Strings(names)

	//Nesse for ele vai pegar o slice ordenado
	//Ele vai pegar esses nomes que estão no slice(as chaves) e irão guardar em name
	//O indice não será relevante nisso, por isso que tem o _
	for _, name := range names {

		//Irá printar a chave, que estará em name, e irá pegar o valor
		fmt.Println(name, ages[name])
	}

	//map nil
	var notas map[string]int
	fmt.Println(notas == nil) //True
	fmt.Println(len(notas))   //0

	//Lembrando que depois que você declara seu mapa, você não consegue adicionar novas chaves, se não usar o make

	//Verificar se existe a chave no mapa
	age, ok := ages["Alice"]
	if !ok {
		fmt.Println("Alice não existe no mapa")
	}
}
