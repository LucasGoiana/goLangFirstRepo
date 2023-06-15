package main

import "fmt"

func main() {
	idade := 10

	if idade >= 18 {

		fmt.Println("Maior de Idade")

	} else {

		fmt.Println("Menor de Idade")

	}

	if outraIdade := idade; outraIdade >= 18 {

		fmt.Println("Maior de Idade")

	} else {

		fmt.Println("Menor de Idade")

	}

}
