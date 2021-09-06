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
	populacao := make([][]individuo, tamanhoPopulacao + 2)
	linhaAux1 := make([]individuo, tamanhoPopulacao + 2) // F para todos os infectados
	linhaAux2 := make([]individuo, tamanhoPopulacao + 2) // F para os extremos e O para o restante
	

	for i := 0; i < tamanhoPopulacao + 2; i++ {
		linhaAux1[i] = individuo{"F", false, 0}

	}

	for i := 0; i < tamanhoPopulacao + 2; i++ {
		if((i == 0) || (i == tamanhoPopulacao + 1)){
			linhaAux2[i] = individuo{"F", false, 0}

		}else{
			linhaAux2[i] = individuo{"O", false, 0}

		}

	}

	// populacao inicial, todos saudaveis mas nao imunizados
	for i := 0; i < tamanhoPopulacao + 2; i++ {
		if((i == 0) || (i == tamanhoPopulacao + 1)){
			populacao[i] = linhaAux1

		}else{
			populacao[i] = linhaAux2

		}

	}

	return populacao

}

func iniciaInfectado(populacao... []individuo) ([][]individuo){
	linhaInfectado := rand.Intn(tamanhoPopulacao) + 1
	colunaInfectado := rand.Intn(tamanhoPopulacao) + 1
	chanceContaminacao := rand.Intn(100) + 1

	infectado := individuo{"X", false, chanceContaminacao}
	linhaInfectadoAux := make([]individuo, tamanhoPopulacao + 2)

	for i := 0; i < tamanhoPopulacao + 2; i++ {
		if (i != colunaInfectado){
			if((i == 0) || (i == tamanhoPopulacao + 1)){
				linhaInfectadoAux[i] = individuo{"F", false, 0}
	
			}else{
				linhaInfectadoAux[i] = individuo{"O", false, 0}
	
			}

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
			if(individuo.situacao != "F"){
				fmt.Printf("%v ", individuo.situacao)

			}

		}

		fmt.Println("")

	}

}
