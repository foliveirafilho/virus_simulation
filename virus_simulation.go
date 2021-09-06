package main

import (
	"fmt"
	"math/rand"
	"time"
)

type individuo struct {
	situacao           string // X para infectado, O para nao infectado
	imunizado          bool
	chanceContaminacao int // chance que um individuo possui para contaminar o outro

}

var tamanhoPopulacao = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	populacao := iniciaPopulacao() // inicia populacao com todos nao infectados
	imprimePopulacao(populacao...)
	populacao = iniciaInfectado(populacao...) // gera-se um individuo aleatorio infectado
	imprimePopulacao(populacao...)
	

}

func iniciaPopulacao() ([][]individuo) {
	populacao := make([][]individuo, tamanhoPopulacao)
	linha := make([]individuo, tamanhoPopulacao)

	for i := 0; i < tamanhoPopulacao; i++ {
		linha[i] = individuo{"O", false, 0}

	}

	// populacao inicial, todos saudaveis mas nao imunizados
	for i := 0; i < tamanhoPopulacao; i++ {
		populacao[i] = linha

	}

	return populacao

}

func iniciaInfectado(populacao... []individuo) ([][]individuo){
	linhaInfectado := rand.Intn(tamanhoPopulacao)
	colunaInfectado := rand.Intn(tamanhoPopulacao)
	chanceContaminacao := rand.Intn(100) + 1

	infectado := individuo{"X", false, chanceContaminacao}
	linhaInfectadoAux := make([]individuo, tamanhoPopulacao)

	for i := 0; i < tamanhoPopulacao; i++ {
		if (i != colunaInfectado){
			linhaInfectadoAux[i] = individuo{"O", false, 0}

		}else{
			linhaInfectadoAux[i] = infectado

		}

	}

	populacao[linhaInfectado] = linhaInfectadoAux

	return populacao

}

func imprimePopulacao(populacao... []individuo) {
	for _, linha := range populacao {
		for _, individuo := range linha {
			fmt.Printf("%v ", individuo.situacao)

		}

		fmt.Println("")

	}

}
