package main

import (
	"fmt"
	"time"
)

func hello() {
	time.Sleep(1 * time.Second)
	fmt.Println("esto es una gourutine")
}

func main() {
	go hello()
	fmt.Println("Esta es la funcion principal")

	time.Sleep(3 * time.Second)
	fmt.Println("Fin del programa")
}
