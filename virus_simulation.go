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

func main() {
	rand.Seed(time.Now().UnixNano())

	tamanho := 10

	populacao := make([][]individuo, tamanho)
	linha := make([]individuo, tamanho)

	for i := 0; i < tamanho; i++ {
		linha[i] = individuo{"O", false, 0}

	}

	// populacao inicial, todos saudaveis mas nao imunizados
	for i := 0; i < tamanho; i++ {
		populacao[i] = linha

	}

	// gera-se um individuo aleatorio infectado
	linhaInfectado := rand.Intn(tamanho)
	colunaInfectado := rand.Intn(tamanho)
	chanceContaminacao := rand.Intn(100) + 1

	infectado := individuo{"X", false, chanceContaminacao}
	linhaInfectadoAux := make([]individuo, tamanho)

	for i := 0; i < tamanho; i++ {
		if (i != colunaInfectado){
			linhaInfectadoAux[i] = individuo{"O", false, 0}

		}else{
			linhaInfectadoAux[i] = infectado

		}

	}

	populacao[linhaInfectado] = linhaInfectadoAux

	fmt.Printf("linha = %v coluna = %v chance = %v\n", linhaInfectado, colunaInfectado, chanceContaminacao)

	for _, linha := range populacao {
		for _, individuo := range linha {
			fmt.Printf("%v ", individuo.situacao)

		}

		fmt.Println("")

	}

}
