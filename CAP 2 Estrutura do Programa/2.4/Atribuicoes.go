package main

import (
	"fmt"
)

func main() {
	//Atribuir é basicamente você colocar um valor em uma variavel que já existe
	var x int
	x = 5

	//Lembrar que os tipos devem ser compativeis
	var y int
	y = int(3.0)

	fmt.Println(x, y)

	//Atribuições multiplas(tuplas)
	//Atribuir varias variaveis de uma vez
	a, b := 1, 2

	//Swap
	//Basicamente troca de valores
	a, b = b, a

	//Atribuições que retornam varios valores
	//Atribuição multipla, mas mostrando quando se trata de valores retornados de uma função
	c, d := dois_valores()

	fmt.Println(c, d)

	//Exemplo comum
	//Caso ocorra um erro, ele retorna o tipo de erro, caso não, retorna o valor
	//Pode também usar _ para ignorar valores, caso não precise de todos os retornos
	var resultado, err = divisao(10, 0)
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	fmt.Println("Resultado = ", resultado)

	//Exemplo com mapa
	//valor é o valor da chave, e ok vai receber true se essa chave existir
	m := map[string]int{
		"o": 1,
	}
	valor, ok := m["o"]

	if ok {
		fmt.Println(valor)
	}

	//Atribuições e ponteiros
	h := 10
	p := &h
	*p = 15

	//Atribuições e Struct
	type Pessoa struct {
		Nome  string
		Idade int
	}

	u := Pessoa{"Inacio", 20}
	u.Idade = 30
}

func dois_valores() (int, int) {
	return 2, 3
}

func divisao(e, f int) (int, error) {
	if f == 0 {
		return 0, fmt.Errorf("Divisão por zero")
	}
	return e / f, nil
}
