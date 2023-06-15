package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	email     string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	estudante1 := estudante{pessoa{"Lucas", "Goiana Malicia", 17, "lucasgoiaanam@hotmail.com"}, "Análise de Desenvolvimento de Sistemas", "FIAP"}

	fmt.Println(estudante1.email)
}
