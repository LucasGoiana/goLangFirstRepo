package main

import "fmt"

func soma(numeros ...int) (soma int) {

	for _, numero := range numeros {
		soma += numero
	}
	return
}
func main() {
	fmt.Println("Funções Váriaticas")
	fmt.Println("A soma desses números é: ",
		soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14))

}
