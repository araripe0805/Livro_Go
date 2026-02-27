package main

//Supondo tenha:
//nums := []int{1,2,3,4}
//E deseja calcular o dobro de cada elemento da lista e somar tudo no final
//2+4+6+8 = 20

//Versão normal(sequencial)
func somaDobros(nums []int) int {
	total := 0

	for _, n := range nums {
		dobro := n * 2
		total += dobro
	}
	return total
}

//1 -> calcula
//2 -> calcula
//3 -> calcula
//4 -> calcula
//executa um por vez

//Ideia de paralelismo
//Cada calculo é independente
//O calculo 1 não depende do 2
//Então pode fazer
//1 -> calcula; 2 -> calcula; 3 -> calcula; 4 -> calcula
//Ao mesmo tempo

//Primeira tentativa de goroutine
//for _, n := range nums {
//	go fmt.Println(n * 2)
//}
//O problema dessa solução é o fato que o main pode terminar antes das goroutines rodarem
//Então precisa esperar as goroutines terminarem

//Se sabe-se que irá receber 4 valores
//precisa-se de 4 goroutines para enviar os resultados
//E receber 4 valores

//Versão em paralela(com quantidade conhecida)
func somaDosResultadosDobros(nums []int) int {
	//Cria um canal
	ch := make(chan int)

	//Cria goroutine a cada index de nums
	for _, n := range nums {
		go func(num int) {
			ch <- num * 2
		}(n)
	}
	//Goroutine A -> envia 2;
	//Goroutine B -> envia 4;
	//Goroutine C -> envia 6;
	//Goroutine D -> envia 8;
	//A ordem pode variar

	//Receber resultados
	total := 0
	for i := 0; i < len(nums); i++ {
		//Bloqueia até receber valores
		total += <-ch
	}
	//Recebe 2
	//Recebe 4
	//Recebe 6
	//Recebe 8
	//A ordem não importa
	//O que importa é receber 4 valores
	return total
}

//Para nums := []int{1,2,3,4}
//LOOP cria 4 goroutines ->
//Todas começam ao mesmo tempo ->
//Cada uma envia resultado no canal ->
//Loop principal recebe 4 valores ->
//Termina
