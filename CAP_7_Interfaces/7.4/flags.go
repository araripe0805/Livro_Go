package main

import (
	"flag"
	"fmt"
	"time"
)

// No pacote flag, existe a interface Value
// Se um tipo implementa esses 2 métodos, ele pode virar uma flag personalizada
// Set(string) recebe um texto digitado pelo usuario e converte para o tipo correto
// String() define como o valor será exibido
type Value interface {
	String() string
	Set(string) error
}

//flag.Duration é uma FUNÇÃO, NÃO UM TIPO
//func Duration(name string, value time.Duration, usage string) *time.Duration
//Recebe um valor default do tipo time.Duration
//Cria uma variavel interna do tipo time.Duration
//Registra essa variavel como uma flag
//Retorna um *time.Duration (ponteiro)

// Mesmo que você não escreva "time.Duration" explicitamente, a variavel period é do tipo *time.Duration
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	//flag.Parse() faz o seguinte:
	//1) Lê os argumentos digitados na linha de comando. Exemplo: ./sleep -period 2m30s
	//2) Encontra a flag "-period"
	//3) Pega o texto "2m30s"
	//4) Chama internamente o método Set("2m30s") da implementação interna da flag de duração
	//5) Esse Set usa: time.ParseDuration("2m30s")
	//6) O valor convertido é armazenado na variavel period
	flag.Parse()

	//*period desreferencia o ponteiro. O tipo real aqui é time.Duration
	//O fmt usa automaticamente o método: (time.Duration).String()
	//Por isso ele imprime: 2m30s e não o número de nanossegundos.
	fmt.Printf("Sleeping for %v...\n", *period)

	//time.Sleep recebe um time.Duration. Como *period é um time.Duration, ele pode ser usado diretamente aqui
	time.Sleep(*period)
}

//Execução:
//./sleep -period 2m30s
//Fluxo interno real:
//Usuário digita:
//-period 2m30s ->
//flag.Parse() ->
//Set("2m30s") ->
//time.ParseDuration("2m30s") ->
//Converte para nanossegundos ->
//Armazena dentro de period ->
//fmt usa period.String() ->
//Imprime: 2m30s

//Criando a propria flag
//Exemplo: Celsius
//Criar um tipo, Implementar Set, implementar String, registra com flag.var

type Celsius float64
type Fahrenheit float64

func conversao(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Agora Celsius implementa fmt.Stringer
func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var value float64
	var unit string

	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = conversao(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

// Esse erro é porque tem 2 funções main, no mesmo codigo, mas para não criar outro file, fiz aqui mesmo
func main() {
	flag.Parse()
	fmt.Println(*temp)
}
