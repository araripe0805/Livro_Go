package main

import (
	"fmt"
)

// Go não possui classes
// O que substitui classes: struct + metodos
// Struct = atributos
// Metodos = comportamento
// Atributos são os campos da struct, se for maiusculo é publico, minusculo é privado
type Pessoa struct {
	Nome   string
	Idade  int
	altura int
}

func (p Pessoa) Falar() {
	fmt.Println("Oi")
}

func (p Pessoa) Saudacao() string {
	return "Olá, eu sou " + p.Nome
}

// Forma de tu acessar um atributo privado
func (p Pessoa) Altura() int {
	return p.altura
}

// Objeto em Go, é basicamente uma instancia de um struct
var p = Pessoa{
	Nome:  "Inacio",
	Idade: 20,
}

func main() {
	fmt.Println(p.Saudacao())
}

//Go não possui herança, apenas composição(Um struct dentro de outro)
//Não é possivel mudar algum valor já inicializado nos atributos de um struct num metodo, apenas se usar ponteiros
