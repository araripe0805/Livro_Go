package main

import (
	"net"
)

//Uma rede é basicamente uma comunicação entre dois computadores
//Navegador conversa com o servidor do Google
//Computador conversa com o banco de dados
//Essa conversa acontece enviando dados na forma binaria

//Quando se vê "localhost: 8000". Significa Conecte nesse computador
//localhost é um nome especial que aponta para o endereço

//Conceito de porta
//Imaginando que o computador é um predio
//Endereço IP é o endereço do predio
//Minha porta é o numero do apartamento
//IP:Porta
//Predio:Apartamento
//localhost:8000
//Esse conceito de porta é usado pois o computador pode estar rodando varios programas de rede ao mesmo tempo
//Exemplo:
//Navegador -> 443
//Banco de dados -> 5432
//Servidor Go -> 8000

//TCP é um modo confiavel de enviar dados entre dois computadores
//Os dados chegam
//Chegam na ordem correta
//Se falhar , ele tenta reenviar

//socket é como um telefone aberto entre dois programas
//Quando ambos os programas querem conversar
//Um programa abre a porta e espera ligação
//Outro programa liga
//Uma conexão é criada
//Eles trocam dados
//Esse canal aberto é o socket
//Cliente -> socket TCP -> Servidor

// Criando um servidor
// Significa: Criar um programa  que fique esperando alguém se conectar a ele
// Para o computador virar um servidor, você só precisa escolher uma porta, dizer ao SO "Quero escutar conexões nessa porta", e esperar alguém tentar conectar
func main() {
	//net é um pacote do Go, que trabalha com rede
	//Ele consegue criar conexões, abrir portas, enviar dados, receber dados
	//O Listen significa "começar a escutar conexões"
	//Quando chama net.Listen, o Go pede ao SO: "Reserve a porta 8000 para mim e me avise quando alguem tentar conectar"
	//listener é um objeto que representa: "A porta aberta esperando conexão"
	listener, err := net.Listen("tcp", "localhost:8000")
	//O SO faz um socket
	//Depois associa à porta 8000
	//Depois marca como "Listening"(modo escuta)
	//Depois reserva essa porta para seu programa
	//Se outro programa já estiver usando a porta 8000 vai dar erro

	//Isso significa: "Espere até alguem tentar se conectar"
	//Esse comando bloqueia
	//bloquear significa: O programa fica parado aqui até um cliente conectar
	conn, err := listener.Accept()

	//Quando um cliente executa
	net.Dial("tcp", "localhost:8000")
	//significa: Tente ligar para locahost na porta 8000
	//Se o servidor estiver rodando: A conexão é criada, Accept() retorna, e o servidor recebe um objeto conn

	//conn representa: A conexão aberta entre cliente e servidor
	//É o "telefone aberto"
	//Pode-se ler dados dela, escrever dados nela, fecha ela
}
