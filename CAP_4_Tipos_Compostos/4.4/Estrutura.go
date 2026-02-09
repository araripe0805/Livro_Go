package main

import (
	"fmt"
)

//Struct é um tipo composto que agrupa varios valores relacionados em uma unica entidade
//Em vez de varias variaveis soltas, juntam toda em um "pacote"

// Pessoa é o tipo
// Nome e Idade são campos
type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	//Criando uma variavel struct
	var p Pessoa
	p.Nome = "Inacio"
	p.Idade = 20

	//versão literal
	a := Pessoa{
		Nome:  "Inacio",
		Idade: 20,
	}

	//Acesso aos campos
	fmt.Println(a.Nome)

	//Struct + ponteiro
	po := &a
	po.Idade = 22

	//Função com struct
	i := Pessoa{"Jose", 18}
	aniversario(&i)

	//Struct pode ser formado por struct
	type Endereco struct {
		endereco string
	}

	type cidade struct {
		Nome     string
		Endereco Endereco
	}

	c := cidade{
		Nome: "Fortaleza",
		Endereco: Endereco{
			endereco: "Beira Mar",
		},
	}

	//Struct recursiva
	//Uso de ponteiros
	type No struct {
		Valor int
		Prox  *No
	}
}

func aniversario(i *Pessoa) {
	i.Idade++
}
