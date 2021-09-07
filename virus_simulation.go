package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/inancgumus/screen"

)

type individuo struct {
	situacao           string // X para infectado, O para nao infectado
	chanceContaminacao int // chance que um individuo possui para contaminar o outro

}

var tamanhoPopulacao = 5

func main() {
	rand.Seed(time.Now().UnixNano())

	populacao := iniciaPopulacao() // inicia populacao com todos nao infectados
	imprimePopulacao(populacao...)

	populacao, linhaInfectado, colunaInfectado := iniciaInfectado(populacao...) // gera-se um individuo aleatorio infectado
	imprimePopulacao(populacao...)

	limpaTela()
	populacao = infectaPopulacao(linhaInfectado, colunaInfectado, populacao...)

	imprimePopulacao(populacao...)

}

func iniciaPopulacao() ([][]individuo) {
	populacao := make([][]individuo, tamanhoPopulacao + 2)
	/*linhaAux1 := make([]individuo, tamanhoPopulacao + 2) // F para todos os infectados
	linhaAux2 := make([]individuo, tamanhoPopulacao + 2) // F para os extremos e O para o restante*/
	
	for i := 0; i < tamanhoPopulacao + 2; i++ {
		linha := make([]individuo, tamanhoPopulacao + 2)

		for j := 0; j < tamanhoPopulacao + 2; j++ {
			if((i == 0) || (i == tamanhoPopulacao + 1) || (j == 0) || (j == tamanhoPopulacao + 1)){
				linha[j] = individuo{"F", 0}

			}else{
				linha[j] = individuo{"O", 0}

			}

		}

		populacao[i] = linha

	}

	/*for i := 0; i < tamanhoPopulacao + 2; i++ {
		linhaAux1[i] = individuo{"F", 0}

	}

	for i := 0; i < tamanhoPopulacao + 2; i++ {
		if((i == 0) || (i == tamanhoPopulacao + 1)){
			linhaAux2[i] = individuo{"F", 0}

		}else{
			linhaAux2[i] = individuo{"O", 0}

		}

	}

	// populacao inicial, todos saudaveis mas com chance de serem contaminados
	for i := 0; i < tamanhoPopulacao + 2; i++ {
		if((i == 0) || (i == tamanhoPopulacao + 1)){
			populacao[i] = linhaAux1

		}else{
			populacao[i] = linhaAux2

		}

	}*/

	return populacao

}

func iniciaInfectado(populacao... []individuo) ([][]individuo, int, int){
	linhaInfectado := rand.Intn(tamanhoPopulacao) + 1
	colunaInfectado := rand.Intn(tamanhoPopulacao) + 1
	chanceContaminacao := rand.Intn(100) + 1

	populacao[linhaInfectado][colunaInfectado].situacao = "X"
	populacao[linhaInfectado][colunaInfectado].chanceContaminacao = chanceContaminacao

	/*infectado := individuo{"X", chanceContaminacao}
	linhaInfectadoAux := make([]individuo, tamanhoPopulacao + 2)

	for i := 0; i < tamanhoPopulacao + 2; i++ {
		if (i != colunaInfectado){
			if((i == 0) || (i == tamanhoPopulacao + 1)){
				linhaInfectadoAux[i] = individuo{"F", 0}
	
			}else{
				linhaInfectadoAux[i] = individuo{"O", 0}
	
			}

		}else{
			linhaInfectadoAux[i] = infectado

		}

	}

	populacao[linhaInfectado] = linhaInfectadoAux*/

	fmt.Println("Primeiro caso do vírus identificado! Chance de contaminação:", chanceContaminacao, "%")
	time.Sleep(3 * time.Second)

	return populacao, linhaInfectado, colunaInfectado

}

func infectaPopulacao(linhaInfectado, colunaInfectado int, populacao... []individuo) ([][]individuo){
	if((linhaInfectado > 0) && (linhaInfectado <= tamanhoPopulacao) && (colunaInfectado > 0) && (colunaInfectado <= tamanhoPopulacao)){
		chanceDeSerInfectado := rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado][colunaInfectado + 1].situacao == "O")){
			populacao[linhaInfectado][colunaInfectado + 1].situacao = "X"
			populacao[linhaInfectado][colunaInfectado + 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v, %v)!\n", linhaInfectado, colunaInfectado, linhaInfectado, colunaInfectado + 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado, colunaInfectado + 1, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado + 1][colunaInfectado + 1].situacao == "O")){
			populacao[linhaInfectado + 1][colunaInfectado + 1].situacao = "X"
			populacao[linhaInfectado + 1][colunaInfectado + 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v, %v)!\n", linhaInfectado, colunaInfectado, linhaInfectado + 1, colunaInfectado + 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado + 1, colunaInfectado + 1, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado + 1][colunaInfectado].situacao == "O")){
			
			populacao[linhaInfectado + 1][colunaInfectado].situacao = "X"
			populacao[linhaInfectado + 1][colunaInfectado].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v, %v)!\n", linhaInfectado, colunaInfectado, linhaInfectado + 1, colunaInfectado)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado + 1, colunaInfectado, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado + 1][colunaInfectado - 1].situacao == "O")){
			
			populacao[linhaInfectado + 1][colunaInfectado - 1].situacao = "X"
			populacao[linhaInfectado + 1][colunaInfectado - 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v, %v)!\n", linhaInfectado, colunaInfectado, linhaInfectado + 1, colunaInfectado - 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado + 1, colunaInfectado - 1, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado][colunaInfectado - 1].situacao == "O")){
			
			populacao[linhaInfectado][colunaInfectado - 1].situacao = "X"
			populacao[linhaInfectado][colunaInfectado - 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v, %v)!\n", linhaInfectado, colunaInfectado, linhaInfectado, colunaInfectado - 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado, colunaInfectado - 1, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado - 1][colunaInfectado - 1].situacao == "O")){
			
			populacao[linhaInfectado - 1][colunaInfectado - 1].situacao = "X"
			populacao[linhaInfectado - 1][colunaInfectado - 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v, %v)!\n", linhaInfectado, colunaInfectado, linhaInfectado - 1, colunaInfectado - 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado - 1, colunaInfectado - 1, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado - 1][colunaInfectado].situacao == "O")){
			
			populacao[linhaInfectado - 1][colunaInfectado].situacao = "X"
			populacao[linhaInfectado - 1][colunaInfectado].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v, %v)!\n", linhaInfectado, colunaInfectado, linhaInfectado - 1, colunaInfectado)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado - 1, colunaInfectado, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado - 1][colunaInfectado + 1].situacao == "O")){
			
			populacao[linhaInfectado - 1][colunaInfectado + 1].situacao = "X"
			populacao[linhaInfectado - 1][colunaInfectado + 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v, %v)!\n", linhaInfectado, colunaInfectado, linhaInfectado - 1, colunaInfectado + 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado - 1, colunaInfectado + 1, populacao...)
			imprimePopulacao(populacao...)

		}

		/*fmt.Printf("linha = %v coluna = %v\n", linhaInfectado, colunaInfectado)
		time.Sleep(time.Second)*/

	}

	return populacao

}

func imprimePopulacao(populacao... []individuo) {
	limpaTela()

	fmt.Println("Acompanhe o desenvolvimento do vírus em tempo real:")
	for _, linha := range populacao {
		for _, individuo := range linha {
			if(individuo.situacao != "F"){
				fmt.Printf("%v ", individuo.situacao)

			}

		}

		fmt.Println("")

	}

	time.Sleep(time.Second)

}

func limpaTela(){
	screen.Clear()
	screen.MoveTopLeft()

}