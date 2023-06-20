package main

import "fmt"

func inverterSinal(n1 int) int {
	return n1 * -1
}

func inverterSinalComPonteiro(n1 *int) {

	*n1 = *n1 * -1

}

func main() {
	n1 := 20
	n1Invertido := inverterSinal(n1)

	fmt.Println(n1, n1Invertido)

	inverterSinalComPonteiro(&n1)
	fmt.Println(n1)

}
