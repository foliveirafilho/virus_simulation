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

var tamanhoPopulacao = 3

func main() {
	rand.Seed(time.Now().UnixNano())

	populacao := iniciaPopulacao() // inicia populacao com todos nao infectados
	imprimePopulacao(populacao...)
	populacao, linhaInfectado, colunaInfectado := iniciaInfectado(populacao...) // gera-se um individuo aleatorio infectado
	imprimePopulacao(populacao...)

	fmt.Printf("linha = %v coluna = %v\n", linhaInfectado, colunaInfectado)

	//limpaTela()
	infectaPopulacao(linhaInfectado, colunaInfectado, "X", "O", populacao...)

	imprimePopulacao(populacao...)

}

func iniciaPopulacao() ([][]individuo) {
	populacao := make([][]individuo, tamanhoPopulacao + 2)
	linhaAux1 := make([]individuo, tamanhoPopulacao + 2) // F para todos os infectados
	linhaAux2 := make([]individuo, tamanhoPopulacao + 2) // F para os extremos e O para o restante
	

	for i := 0; i < tamanhoPopulacao + 2; i++ {
		linhaAux1[i] = individuo{"F", 0}

	}

	for i := 0; i < tamanhoPopulacao + 2; i++ {
		if((i == 0) || (i == tamanhoPopulacao + 1)){
			linhaAux2[i] = individuo{"F", 0}

		}else{
			linhaAux2[i] = individuo{"O", 0}

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

func iniciaInfectado(populacao... []individuo) ([][]individuo, int, int){
	linhaInfectado := rand.Intn(tamanhoPopulacao) + 1
	colunaInfectado := rand.Intn(tamanhoPopulacao) + 1
	chanceContaminacao := rand.Intn(100) + 1

	infectado := individuo{"X", chanceContaminacao}
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

	populacao[linhaInfectado] = linhaInfectadoAux

	return populacao, linhaInfectado, colunaInfectado

}

func infectaPopulacao(linhaInfectado, colunaInfectado int, situacaoAnterior, situacaoAtual string, populacao... []individuo) {
	if((linhaInfectado > 0) && (linhaInfectado <= tamanhoPopulacao) && (colunaInfectado > 0) && (colunaInfectado <= tamanhoPopulacao) && 
	(populacao[linhaInfectado][colunaInfectado].situacao == situacaoAnterior) && (populacao[linhaInfectado][colunaInfectado].situacao != situacaoAtual)){
		chanceDeSerInfectado := rand.Intn(100) + 1

		fmt.Printf("ser = %v passar = %v\n", chanceDeSerInfectado, populacao[linhaInfectado][colunaInfectado].chanceContaminacao)

		populacao[linhaInfectado][colunaInfectado].situacao = "X"
		populacao[linhaInfectado][colunaInfectado].chanceContaminacao = chanceDeSerInfectado

		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado][colunaInfectado + 1].situacao == "O")){
			infectaPopulacao(linhaInfectado, colunaInfectado + 1, situacaoAtual, situacaoAnterior, populacao...)
			imprimePopulacao(populacao...)

		}

		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado + 1][colunaInfectado + 1].situacao == "O")){
			infectaPopulacao(linhaInfectado + 1, colunaInfectado + 1, situacaoAtual, situacaoAnterior, populacao...)
			imprimePopulacao(populacao...)

		}

		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado + 1][colunaInfectado].situacao == "O")){
			infectaPopulacao(linhaInfectado + 1, colunaInfectado, situacaoAtual, situacaoAnterior, populacao...)
			imprimePopulacao(populacao...)

		}

		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado + 1][colunaInfectado - 1].situacao == "O")){
			infectaPopulacao(linhaInfectado + 1, colunaInfectado - 1, situacaoAtual, situacaoAnterior, populacao...)
			imprimePopulacao(populacao...)

		}

		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado][colunaInfectado - 1].situacao == "O")){
			infectaPopulacao(linhaInfectado, colunaInfectado - 1, situacaoAtual, situacaoAnterior, populacao...)
			imprimePopulacao(populacao...)

		}

		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado - 1][colunaInfectado - 1].situacao == "O")){
			infectaPopulacao(linhaInfectado - 1, colunaInfectado - 1, situacaoAtual, situacaoAnterior, populacao...)
			imprimePopulacao(populacao...)

		}

		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado - 1][colunaInfectado].situacao == "O")){
			infectaPopulacao(linhaInfectado - 1, colunaInfectado, situacaoAtual, situacaoAnterior, populacao...)
			imprimePopulacao(populacao...)

		}

		if((chanceDeSerInfectado <= populacao[linhaInfectado][colunaInfectado].chanceContaminacao) && (populacao[linhaInfectado - 1][colunaInfectado + 1].situacao == "O")){
			infectaPopulacao(linhaInfectado - 1, colunaInfectado + 1, situacaoAtual, situacaoAnterior, populacao...)
			imprimePopulacao(populacao...)

		}

		/*fmt.Printf("linha = %v coluna = %v\n", linhaInfectado, colunaInfectado)
		time.Sleep(time.Second)*/

	}

}

func imprimePopulacao(populacao... []individuo) {
	//limpaTela()

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