package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media == 6 {
		panic(media)
	}

	return media > 6
}

func main() {

	mediaBoolean := alunoEstaAprovado(6, 6)

	fmt.Println(mediaBoolean)
}
