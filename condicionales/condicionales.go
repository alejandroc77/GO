package main

import "fmt"

func main() {
	x := 17
	y := 21

	if x > y {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("Eres mayor de edad")
	}

	temperature()
	aninados()

}

func temperature() {
	temperatura := 21

	if temperatura < 1 {
		fmt.Println("Exceso de calor, !Epoca de verano")
	} else {
		fmt.Println("Epoca invernal")
	}
}

func aninados() {
	num := 12

	if num >= 3 {
		fmt.Println("El numero es mayor al numero dado")

		if num >= 4 {
			fmt.Println("El numero es mayor")
		}
	} else {
		fmt.Println("el numero no es mayor al dado")
	}
}
