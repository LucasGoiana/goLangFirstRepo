package main

import "fmt"

func main() {
	var variavel1 string = "banana"
	variavel2 := "Testei"

	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Meu Deus", "do Céu"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante1"

	fmt.Println(constante1)

	//Invertando valor
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
