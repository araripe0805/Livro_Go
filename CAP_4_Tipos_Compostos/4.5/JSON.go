package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//JSON é texto, não estrutura
	//Usada para enviar, receber, salvar dados e trocar dados entre sistemas
	//JSON é uma forma padrão de escrever dados em texto para que diferentes sistemas consigam se entender

	{
  		"name": "Inacio",
  		"age": 20,
  		"skills": ["Go", "Java", "Git"],
  		"active": true
	}


	//Marshaling = transforma um valor Go em JSON
	p := Person{Name: "Inacio",
		Age: 20}

	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))// {"Name":"Inacio","Age":20}

	//MarshalIndent (JSON bonito)
	//Marshal normal é compacto
	//Para humanos, usa-se mais o MarshalIdent
	//"" prefixo de linha
	//"  " indentação por nivel
	data2, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(data2)) //{
  							   //  "Name": "Inacio",
  							   //  "Age": 20
							   //}

	//Field tags(json:"..." e omitempty)
	//Permite controlar como o JSON será gerado
	
	//Struct com tags
	type Movie struct{
		Title string
		Year  int `json:"released"`
		Color bool `json:"color,omitempty"`
		Actors []string
	}

	//json:"released"

	//Troca o nome do campo no JSON
	//Year int json:"released"

	//Em Go
	//Year: 1942

	//JSON
	//"released": 1942

	//omitempty não gera o campo se ele estiver vazio(nil,"",0...)
	//Movie {
	//	Title: "Casablanca",
	//	Year: 1942,
	//	Color: false, -> omitido
	//}

	//Unmarshaling (JSON -> Go)
	//quer transformar isso em struct, slice ou map em Go

	//Json
	jsonData := []byte(`
	{
		"Name": "Inacio",
		"Age": 20
	}`)

	var q Person

	//JSON tem que ser valido(primeiro parametro)
	//Segundo parametro tem que ser ponteiro
	err = json.Unmarshal(jsonData, &q)
	if err != nil {
		panic(err)
	}
	fmt.Println(q.Name, q.Age)

	//Unmarshal lê o texto, procura chaves com o mesmo nome e escreve nos campos

	//Decodificar APENAS parte do JSON
	//JSON grande
	{
		"id": 10,
  		"name": "Inacio",
  		"age": 20,
  		"password": "123456",
  		"token": "abc",
  		"admin": true
	}

	//Você não é obrigado a criar um struct com tudo isso.
	type User struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}

	//Unmarshal
	var u User
	json.Unmarshal(data, &u)
	//name e age são preenchidos
	//resto ignorado

	//Fluxo real da web
	//Servidor -> http -> JSON -> Go
	//1. Você faz um http.Get
	//2. O servidor responde com texto JSON
	//3. Esse texto vem em resp.Body
	//4. Você decodifica esse JSON
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
