package main

import (
	"fmt"
	"math/rand"
)

func main() {

	f := func(txt string) string {
		return "oi"
	}

	fmt.Println(f("a"))
	n1 := rand.Float32()
	n2 := rand.Float32()

	fmt.Sprintln("Hi", "a")
	fmt.Sprintln(n1, n2)
	//fmt.Println("A soma de", n1, "+", n2, "é", somar(n1, n2))

	//fmt.Println("A Subtração de", n1, "-", n2, "é", subtrair(n1, n2))

	fmt.Println(dividirEmultiplicar("asdsa", n1, n2))

}

func somar(n1 int, n2 int) int {
	soma := n1 + n2
	return soma
}

func subtrair(n1 int, n2 int) int {
	subtracao := n1 - n2
	return subtracao
}

func dividirEmultiplicar(operacaoAritmetica string, n1, n2 float32) (string, float32) {

	if operacaoAritmetica == "dividir" {
		divisao := n1 / n2
		mensagem := "o resultado é"
		return mensagem, divisao
	}

	multiplicacao := n1 * n2
	mensagem := "o resultado da multiplicação é"
	return mensagem, multiplicacao

}
