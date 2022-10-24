package main

import "fmt"

func familyName(name string, age int) {
	fmt.Println("hello", name, "tengo", age)
}

func myFunction(x int, y int) (result int) { // Valores devueltos con nombres
	result = x + y
	return
}

func myFunction2(x int, y int) int { // aqui se ve como se hace un retorno mas limpio
	return x + y
}

func calculo(a, b int) (area int) {
	parametro := 2 * (a + b)
	fmt.Printf("El resultado del parametro es %v \n", parametro)

	area = a * b
	return
}

func main() {
	familyName("jordy", 21)
	familyName("sandra", 12)
	fmt.Println("Area", calculo(13, 4))
}

//retornos de funciones
