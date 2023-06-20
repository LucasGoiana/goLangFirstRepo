package main

import "fmt"

func main() {
	fmt.Println("Função Anonima")

	retorno := func(texto string) (retorno string) {
		retorno = fmt.Sprintf("Recebido -> %s", texto)
		return
	}("Olá Mundo")

	fmt.Println(retorno)
}
