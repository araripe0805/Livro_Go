package main

import (
	"fmt"
	"sort"
	"time"
)

// O pacote sort permite ordenar qualquer sequencia, desde que ela implemente a interface
// Esses metodos que a interface exige são o: Len() -> tamanho; Less(i,j) -> como comparar dois elementos; Swap(i,j) -> Como trocar dois elementos
// Isso é uma forte arma do Go, o algoritmo pode ser totalmente generico, e quem define tudo é quem implementa sort.Interface
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//Exemplo: Ordenação de strings

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	names := []string{"Carlos", "Ana", "Bruno"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)
}

//Esse tipo de coisa é tão comum que existe atalhos:
//sort.Strings(names)
//sort.Ints(values)
//sort.IntsAreSorted(values)

//Exemplo 2: Playlist de musicas
//Struct
type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

//Criando a playlist
//Eu uso *Track[], ou seja, um slice de ponteiros
//Isso é ajuda muito, pois como o algoritmo de ordenação chama Swap muitas vezes, se fosse um slice comum, ele ia constantemente ficar criando copias do structs para pequenas alterações, e deixando a execução mais lenta
var tracks = []*Track{{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},}


//Ordenando por artista
type byArtist []*Track

func (b byArtist) Len() int {
	return len(b)
}

func (b byArtist) Less(i,j int) bool {
	return b[i].Artist < b[j].Artist
}

func (b byArtist) Swap(i,j int) {
	b[i],b[j] = b[j],b[i]
}

//Agora pode ser feita uma ordenação pela ordem de autores, ou seja, é possivel decidir a forma de ordenação
sort.Sort(byArtist(tracks))

//Ordem reversa
//Não preciso criar um novo tipo
//O pacote fornece:
sort.Sort(sort.Reverse(byArtist(tracks)))
//Internamente existe um tipo reverse, e ele implementa o metodo Less da seguinte forma:
func (r Reverse) Less(i,j int) bool{
	return r.Interface.Less(j,i)
}
//ele apenas inverte os indices, e reutiliza Len e Swap automaticamente

//É perceptivel que o Len, e o Swap sempre são iguais, apenas o Less que muda, então para solucionar a repetição de codigo:
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (c customSort) Len() int {
	return len(c.t)
}

func (c customSort) Less(i, j int) bool {
	return c.less(c.t[i], c.t[j])
}

func (c customSort) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

//Agora pode passar qualquer regra
sort.Sort(customSort{
	tracks,
	func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		return false
	},
})