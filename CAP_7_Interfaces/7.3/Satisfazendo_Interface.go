package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
)

//Um tipo satisfaz uma interface se ele possui todos os métodos que a interface exige
type Writer interface{
	Write(p []byte) (n int, err error)
}

type MyFile struct{}

func(m MyFile) Write(p []byte) (n int,err error) {
	fmt.Println(string(p))
	return len(p), nil
}

//MyFile satisfaz porque possui o metodo Write
//Logo, não a problema em atribuir uma variavel do tipo da interface, com um struct que possui o metodo listado na interface
var w Writer
w = MyFile{}//ok

//*os.File tem métodos Read, Write e Close
//Por isso ele satisfaz io.Reader, io.Writer, io.Closer, io.ReadWriter, io.ReadWriteCloser
//Porque os.Stdout é do tipo *os.File, e *os.File tem método Write
var a io.Writer
a = os.Stdout//ok

//*bytes.Buffer tem Read e Write, mas não tem Close
var b io.Writer
b = new(bytes.Buffer) //ok

//Mas
//Isso ocorre porque rwc é do tipo que possui metodos Write, Read e Close
//Mas está tentando receber uma interface que só possui Write e Read, não possui o close, logo ele não satisfaz a interface do tipo rwc
var rwc io.ReadWriteCloser
rwc = new(bytes.Buffer) //Erro

//Resumindo, uma variavel do tipo de uma interface só pode receber um tipo que satisfaça a interface

//Interface recebendo interface
//Sabemos que io.Writer só tem Write e io.ReadWriteCloser tem entre um de seus métodos o Write
var c io.Writer
var rwc2 io.ReadWriteCloser

//Logo
//Isso ocorre pois c só tem Write como método, então basta qualquer outra variavel ter um método Write, e dentre seus vários métodos, rwc2 tem
c = rwc2 //ok

//Mas
//Isso da erro porque rwc2 possui os metodos Write,Read, Close, e estou tentandoa atribuir uma outra interface que só possui Write, logo não satisfaz, porque faltará as outras 2
rwc2 = c //Erro

//Regra: Interface maior pode virar a menor, já a menor não pode virar a maior	

//Ponteiro X Valor
type IntSet struct{}

//*IntSet tem o método String
//IntSet não tem o método String
//Isso ocorre pelo receptor do método, o metodo não pertence a IntSet
func (*IntSet) String() string {
	return "IntSet"
}

//Internamente o compilador transforma em (&s).String()
var s IntSet
s.String() //Ok

//Mas quando se trata de interface
//Esse erro acontece pois somente *IntSet possui o metodo
var _ fmt.Stringer = &s // ok
var _ fmt.Stringer = s // erro

//exemplo mais claro
type Exemplo struct{}

func (Exemplo) valor()
func (*Exemplo) ponteiro()

//Exemplo: valor()
//*Exemplo: valor(), ponteiro()

//Mesmo o os.Stdout possuindo varios outros metodos, d não irá conseguir usufruir deles, pois sua interface possui apenas o metodo Write
var d io.Writer
d = os.Stdout

//Interface vazia, ela não possui nenhum método, então qualquer coisa pode ser atribuida
var any interface{}
any = 10
any = "hello"
any = true
any = new(bytes.Buffer)
