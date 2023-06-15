package main

import (
	"errors"
	"fmt"
)

func main() {

	numero := 1
	fmt.Println(numero)

	var erro error = errors.New("Erro interno")

	fmt.Println(erro)
}
