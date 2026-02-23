package main

import (
	"fmt"
	"log"
	"net/http"
)

//A base de qualquer servidor HTTP em Go é a interface Handler
type Handler interface{
	ServeHTTP(w ResponseWriter, r *Request)
}

//Essa função recebe um endereço e algo que implemente http.Handler
http.ListenAndServe(address string, h Handler)

//Exemplo: Tipo que implementa ServeHTTP
import(
	"fmt"
	"log"
	"net/http"
)

type MeuHandler struct {}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Olá! Servidor funcionando")
}

func main() {
	handler := MeuHandler{}
	log.Fatal(http.ListenAndServe(":8000", handler))
}

//MeuHandler implementa ServeHTTP
//Portanto ele satisfaz http.Handler
//Pode ser passado para ListenAndServe