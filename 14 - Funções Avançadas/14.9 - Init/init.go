package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando a Init")
	n = 10
}

func main() {
	fmt.Println(n)
}
