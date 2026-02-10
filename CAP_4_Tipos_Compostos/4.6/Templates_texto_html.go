package main

import (
	"html/template"
	"os"
)

func main() {
	//Template = texto normal + buracos que o Go preenche
	//Esses "buracos" ficam entre: {{alguma coisa}}
	//Tudo fora disso é texto literal, e tudo dentro é logica do template
	//O objetivo é separar formato(texto/html) da logica(Go)

	//Estrutura do template
	//1. Criar template
	//t := template.New("nome")

	//2. Fazer o parse do texto do template
	//t = t.Parse(textoDoTemplate)

	//3. Executar o template com dados
	//t.Execute(saida, dados)

	//Forma curta
	//t := template.Must(template.New("nome").Parse(texto))
	//t.Execute(os.Stdout, dados)

	//O . significa dado atual
	//t.Execute(os.Stdout, dados)

	//Se o dado for de um struct
	//type Pessoa struct {
	//	Nome string
	//	Idade int
	//}

	//No template:
	//{{.Nome}}
	//{{.Idade}}

	//Tudo dentro de {{}} é uma ação
	//Ações importantes

	//Mostra valor
	//{{.Nome}}

	//Controle de fluxo
	//{{if .Ativo}} ... {{end}}
	//{{range .Lista}} ... {{end}}

	//Funções
	//{{printf "%.2f" .Preco}}

	//Fora disso -> texto normal
	//Dentro disso -> Comando do template

	//Exemplos
	type Pessoa struct {
		Nome  string
		Idade int
	}

	const tmpl = `Nome: {{.Nome}} Idade: {{.Idade}}`

	t := template.Must(template.New("pessoa").Parse(tmpl))

	p := Pessoa{
		Nome:  "Inácio",
		Idade: 20,
	}

	t.Execute(os.Stdout, p)

	//Execute(..., p) -> . agora é p -> {{.Nome}} -> p.Nome

	//Controle de fluxo em templates
	//Controle de fluxo  serve para decidir se algo aparece ou não, ou repetir bloco de texto

	//If e else
	//{{if condicao}}
	// 	texto se for verdadeiro
	//{{else}}
	//	texto se for falso
	//{{end}}

	type Objeto struct {
		Nome  string
		Ativo bool
	}

	const tmpl2 = `Nome: {{.Nome}} 
	
	{{if .Ativo}}
	Status: Ativo
	{{else}}
	Status: Inativo
	{{end}}`

	t2 := template.Must(template.New("objeto").Parse(tmpl2))

	o := Objeto{Nome: "Tesoura", Ativo: true}

	t2.Execute(os.Stdout, o)

	//Range(repetição)
	//range é o for do template
	//{{range colecao}}
	//	texto repetido
	//{{end}}

	type Teste struct {
		Nome    string
		Hobbies []string
	}

	const tmpl3 = `Hobbies:
	{{range .Hobbies}}
	- {{.}}
	{{end}}
	`

	t3 := template.Must(template.New("teste").Parse(tmpl3))

	t5 := Teste{
		Nome:    "Inácio",
		Hobbies: []string{"Ler", "Jogar"},
	}
	t3.Execute(os.Stdout, t5)
}
