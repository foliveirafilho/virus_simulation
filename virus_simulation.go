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

	fmt.Println("numero =", rand.Intn(100))
	fmt.Println("numero =", rand.Intn(100))

	tamanho := 10

	matriz := make([][]string, tamanho)
	linha := make([]string, tamanho)

	for i := 0; i < tamanho; i++ {
		linha[i] = "O"

	}

	for i := 0; i < tamanho; i++ {
		matriz[i] = linha
	}

	for _, sliceLinha := range matriz {
		for _, valor := range sliceLinha {
			fmt.Printf("%v ", valor)

		}

		fmt.Println("")

	}

}
