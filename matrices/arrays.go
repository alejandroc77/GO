package main

import "fmt"

func main() {
	cars := [3]string{"mazda", "toyota", "bmw"}

	fmt.Println(cars[2])

	cambioValorres()
}

func cambioValorres() {
	number := [3]int{32, 54, 65}

	number[1] = 77 //aqui se cambia el valor haciendo referencia al numero de indice

	fmt.Println(number)
}
