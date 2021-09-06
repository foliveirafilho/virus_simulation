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

	for i := 0; i < tamanho; i++ {
		populacao[i] = linha

	}

	for _, linha := range populacao {
		for _, individuo := range linha {
			fmt.Printf("%v ", individuo.situacao)

		}

		fmt.Println("")

	}

}
