package main

import (
	"errors"
	"fmt"
)

// Em Go, error é uma interface pré-declarada
// Qualquer tipo que tenha o método Error() string automaticamente implementa error
type error interface {
	Error() string
}

//O pacote errors é basicamente isso:

//package errors

type errorString struct {
	text string
}

func New(text string) error {
	return &errorString{text}
}

func (e *errorString) Error() string {
	return e.text
}

//errorString é uma struct que guarda a mensagem
//O método Error() está definido para *errorString
//Como ele possui Error() string, ele satisfaz a interface error
//New retorna um valor do tipo error

// tem mais 2 funções main que vão servir para exemplificar melhor
func main() {
	err := errors.New("algo deu errado")

	fmt.Println(err)         //Chama automaticamente Error()
	fmt.Println(err.Error()) //Mesma coisa

	//algo deu errado
	//algo deu errado
}

// O pacote fmt possui:
func Errorf(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}

func main() {
	nome := "Inácio"
	err2 := fmt.Errorf("Usuario %s não encontrado", nome)
	fmt.Println(err2)
	//Usuario Inácio não encontrado
	//Isso é apenas uma forma formatada de errors.New
}

//error é só uma interface, logo, qualquer tipo pode implementa-la

// Criando um erro estruturado
type ErroBanco struct {
	Codigo   int
	Mensagem string
}

func (e ErroBanco) Error() string {
	return fmt.Sprintf("Erro %d: %s", e.Codigo, e.Mensagem)
}

func main() {
	err3 := ErroBanco{404, "registro não encontrado"}
	fmt.Println(err3)
	//Erro 404: registro não encontrado
}
