package main

import "fmt"

func main() {
	fmt.Println("aaa")

	var array1 [5]int
	array1[1] = 1
	fmt.Println(array1)

	array2 := [5]int{0, 1, 2, 3, 4}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4}
	fmt.Println(array3)

	slice := []int{}

	slice = append(slice, 2)

	fmt.Println(slice)

	slice2 := array2[1:3]

	slice2 = append(slice2, 3)
	fmt.Println(slice2)
}
