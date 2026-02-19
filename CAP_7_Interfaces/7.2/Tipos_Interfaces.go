package main

import (
	"fmt"
)

//Um tipo interface define um conjunto de métodos

// io.Reader
// Representa qualquer coisa de onde se pode ler bytes
// O metodo Read ler os dados da fonte e copia em p
// p []byte é o buffer onde os dados serão colocados
// O retorno possui o n, que é a quantidade de bytes lidos
// err é o erro(caso tenha)
type Reader interface {
	Read(p []byte) (n int, err error)
}

// io.Closer
// Usado quando algo precisa ser usado/liberado
type Closer interface {
	Close() error
}

// io.Writer
// Representa qualquer coisa que pode escrever bytes
// O metodo Write ele vai pegar esses bytes e enviar para o destino
// n retorna quantos bytes foram escritos
// err representa o erro
type Writer interface {
	Write(p []byte) (n int, err error)
}

//Como convenção, normalmente interfaces com um unico metodo terminam com "er"

// Interfaces podem ser combinadas
// Um tipo que satisfaça Reader e Writer
type ReadWriter interface {
	Reader
	Writer
}

//ou
//type ReadWriter interface{
//	Read(p []byte) (n int, err error)
//	Write(p []byte) (n int, err error)
//}

// Exemplo
type FileSimulado struct {
	data []byte
}

func (f *FileSimulado) Read(p []byte) (int, error) {
	n := copy(p, f.data)
	f.data = f.data[n:]
	return n, nil
}

func (f *FileSimulado) Write(p []byte) (int, error) {
	f.data = append(f.data, p...)
	return len(p), nil
}

func (f *FileSimulado) Close() error {
	fmt.Println("Arquivo fechado")
	return nil
}
