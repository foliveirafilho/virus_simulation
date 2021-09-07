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
var qtdeInfectados = 0

func main() {
	rand.Seed(time.Now().UnixNano())

	populacao := iniciaPopulacao() // inicia populacao com todos nao infectados
	imprimePopulacao(populacao...)

	populacao, linhaInfectado, colunaInfectado := iniciaInfectado(populacao...) // gera um individuo aleatorio infectado
	imprimePopulacao(populacao...)

	limpaTela()
	populacao = infectaPopulacao(linhaInfectado, colunaInfectado, populacao...)

	imprimePopulacao(populacao...)

	fmt.Printf("Ao final, a quantidade de infectados foi de %v, ", qtdeInfectados)
	fmt.Println("uma taxa de", (qtdeInfectados / (tamanhoPopulacao * tamanhoPopulacao * 1.0)) * 100,"%!")

}

func iniciaPopulacao() ([][]individuo) {
	populacao := make([][]individuo, tamanhoPopulacao + 2)
	
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

	return populacao

}

func iniciaInfectado(populacao... []individuo) ([][]individuo, int, int){
	linhaInfectado := rand.Intn(tamanhoPopulacao) + 1
	colunaInfectado := rand.Intn(tamanhoPopulacao) + 1
	chanceContaminacao := rand.Intn(100) + 1

	qtdeInfectados++

	populacao[linhaInfectado][colunaInfectado].situacao = "X"
	populacao[linhaInfectado][colunaInfectado].chanceContaminacao = chanceContaminacao

	fmt.Println("Primeiro caso do vírus identificado! Chance de contaminação:", chanceContaminacao, "%")
	time.Sleep(3 * time.Second)

	return populacao, linhaInfectado, colunaInfectado

}

func infectaPopulacao(linhaInfectado, colunaInfectado int, populacao... []individuo) ([][]individuo){
	if((linhaInfectado > 0) && (linhaInfectado <= tamanhoPopulacao) && (colunaInfectado > 0) && (colunaInfectado <= tamanhoPopulacao)){
		chanceDeSerInfectado := rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado][colunaInfectado + 1].situacao == "O")){
			qtdeInfectados++
			populacao[linhaInfectado][colunaInfectado + 1].situacao = "X"
			populacao[linhaInfectado][colunaInfectado + 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v,%v)!\n", linhaInfectado, colunaInfectado, linhaInfectado, colunaInfectado + 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado, colunaInfectado + 1, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado + 1][colunaInfectado + 1].situacao == "O")){
			qtdeInfectados++
			populacao[linhaInfectado + 1][colunaInfectado + 1].situacao = "X"
			populacao[linhaInfectado + 1][colunaInfectado + 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v,%v)!\n", linhaInfectado, colunaInfectado, linhaInfectado + 1, colunaInfectado + 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado + 1, colunaInfectado + 1, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado + 1][colunaInfectado].situacao == "O")){
			qtdeInfectados++
			populacao[linhaInfectado + 1][colunaInfectado].situacao = "X"
			populacao[linhaInfectado + 1][colunaInfectado].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v,%v)!\n", linhaInfectado, colunaInfectado, linhaInfectado + 1, colunaInfectado)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado + 1, colunaInfectado, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado + 1][colunaInfectado - 1].situacao == "O")){
			qtdeInfectados++
			populacao[linhaInfectado + 1][colunaInfectado - 1].situacao = "X"
			populacao[linhaInfectado + 1][colunaInfectado - 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v,%v)!\n", linhaInfectado, colunaInfectado, linhaInfectado + 1, colunaInfectado - 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado + 1, colunaInfectado - 1, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado][colunaInfectado - 1].situacao == "O")){
			qtdeInfectados++
			populacao[linhaInfectado][colunaInfectado - 1].situacao = "X"
			populacao[linhaInfectado][colunaInfectado - 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v,%v)!\n", linhaInfectado, colunaInfectado, linhaInfectado, colunaInfectado - 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado, colunaInfectado - 1, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado - 1][colunaInfectado - 1].situacao == "O")){
			qtdeInfectados++
			populacao[linhaInfectado - 1][colunaInfectado - 1].situacao = "X"
			populacao[linhaInfectado - 1][colunaInfectado - 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v,%v)!\n", linhaInfectado, colunaInfectado, linhaInfectado - 1, colunaInfectado - 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado - 1, colunaInfectado - 1, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado - 1][colunaInfectado].situacao == "O")){
			qtdeInfectados++
			populacao[linhaInfectado - 1][colunaInfectado].situacao = "X"
			populacao[linhaInfectado - 1][colunaInfectado].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v,%v)!\n", linhaInfectado, colunaInfectado, linhaInfectado - 1, colunaInfectado)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado - 1, colunaInfectado, populacao...)
			imprimePopulacao(populacao...)

		}

		chanceDeSerInfectado = rand.Intn(100) + 1
		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado - 1][colunaInfectado + 1].situacao == "O")){
			qtdeInfectados++
			populacao[linhaInfectado - 1][colunaInfectado + 1].situacao = "X"
			populacao[linhaInfectado - 1][colunaInfectado + 1].chanceContaminacao = rand.Intn(100) + 1

			imprimePopulacao(populacao...)
			fmt.Printf("O indivíduo (%v,%v) infectou o indivíduo (%v,%v)!\n", linhaInfectado, colunaInfectado, linhaInfectado - 1, colunaInfectado + 1)
			time.Sleep(2 * time.Second)
			populacao = infectaPopulacao(linhaInfectado - 1, colunaInfectado + 1, populacao...)
			imprimePopulacao(populacao...)

		}

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