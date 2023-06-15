package main

import "fmt"

func main() {

	/*usuario := map[int]string{
		0:      "Pedro",
		1: "Claudin",
	}*/

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Claudin",
	}

	fmt.Println(usuario)

	delete(usuario, "nome")

	fmt.Println(usuario)

	usuario["signo"] = "a"
	fmt.Println(usuario)

}
