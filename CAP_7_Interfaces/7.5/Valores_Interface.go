package main

import "io"

//Quando declara w, ele não guarda só um valor, ele guarda duas coisas ao mesmo tempo: (tipo dinamico, valor dinamico)
//interface = (qual é o tipo real?, qual é o valor?)
var w io.Writer

//Interface nil
//O valor inicial de w é: tipo dinamico = nil ;  valor dinamico = nil

//Agora
//os.Stdout é do tipo *os.File
//Logo, w vira: tipo dinamico = *os.File; valor dinamico = ponteiro para Stdout
//Mesmo que o tipo da variavel seja io.Writer, o real tipo armazenado dentro dela é *os.File	
w = os.Stdout

//Quando chamamos esse metodo:
//O Go olha o tipo dinamico(*os.File), procura o método (Write) e chama((*os.File).Write(...))
//Isso é chamado de Despacho Dinamico
//O metodo é decidido durante o tempo de execução, e não no periodo de compilação
w.Write([]byte(|"hello"))

//Mudando o tipo dinamico
//(tipo dinamico = *bytes.Buffer), (Valor dinamico = ponteiro para buffer)
w = new(bytes.Buffer)

//Assim, quando chamar o metodo, ele irá assim: (*bytes.Buffer).Write(...)
w.Write([]byte("hello"))

//A mesma variavel w, mesmo tipo estatico io.Writer, mas comportamento diferente, uma especie de polimorfismo

//Tipo estatico != tipo dinamico
//Tipo estatico = io.Writer ; Tipo dinamico = *os.File