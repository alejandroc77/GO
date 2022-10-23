package main

import (
	"fmt"
	"os"
	"strconv"
)

// func add(firt int, scond int) int {
// 	return firt + scond
// }

func main() {
	no, err := strconv.Atoi(os.Args[1])
	fmt.Println(no)

	if err != nil {
		fmt.Println(err)
		fmt.Println("no se puedo convertir " + os.Args[1])
	} else {
		fmt.Println(no)
	}

	// num1, _ := strconv.Atoi(os.Args[1])
	// num2, _ := strconv.Atoi(os.Args[2])

	// var sum = add(num1, num2)
	// fmt.Println(sum)
}
