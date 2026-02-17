package main

import (
	"fmt"
)

// Encapsulamento = esconder detalhes internos de um tipo e só permitir acesso apenas através de métodos controlados
// Protege os dados
// Manter invariantes
// Permitir mudar implementação sem quebrar o código
// Inicial maiuscula: Publica
// Inicial minuscula: Privado
type Teste struct {
	Nome  string //Publico
	idade int    //Privado
}

func main() {
	p := Teste{Nome: "inacio", idade: 20}
	fmt.Println(p.Nome)  //ok
	fmt.Println(p.idade) //não daria certo essa linha em outro arquivo
}
