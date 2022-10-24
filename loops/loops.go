package main

import "fmt"

func main() {
	// Example 1
	comida := map[string]string{"manzana": "pera"}

	// for index, element := range comida {
	// 	fmt.Println("index", index, "elemento", element)
	// }

	//	Example 2
	// for key := range comida {
	// 	fmt.Println(key)
	// }

	// // Example 3
	for _, value := range comida {
		fmt.Println(value)
	}
}
