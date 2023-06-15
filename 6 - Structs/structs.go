package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func main() {

	fmt.Println("asdasds")
	var usuario1 usuario
	usuario1.nome, usuario1.idade = "Lucas Goiana Malicia", 27
	fmt.Println(usuario1)

	usuario2 := usuario{"Felipe Goiana Malicia", 999}
	fmt.Println(usuario2)

	usuario3 := usuario{"Suzanira Goiana Malicia", 222}
	fmt.Println(usuario3)

	usuario4 := usuario{nome: "Vanderlei Tortosa Malicia"}
	fmt.Println(usuario4)
}
