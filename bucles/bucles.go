package main

import "fmt"

func main() {
	frutas := [3]string{"maduro", "manzana", "pera"}

	// for index, values := range frutas {
	// 	fmt.Printf("%v \t %v \n", index, values)
	// }
	for _, value := range frutas { // esto es lo mismo de arriba pero se omiten los indices
		fmt.Printf("%v \n", value)
	}
}
