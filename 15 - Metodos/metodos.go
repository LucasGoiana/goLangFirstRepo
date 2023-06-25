package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Println(u.nome)
}

func main() {

	usuario1 := usuario{"Nome", 1}

	fmt.Println(usuario1)
	usuario1.salvar()
}
