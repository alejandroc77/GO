package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GOROTUNES\t\t", runtime.NumGoroutine())

	wg.Add(1)
	foo()
	bar()
	verificar()

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Print("== last foo == ")
		fmt.Println("\tbar:", i)
	}
}

func verificar() { //vrificando tipos de datos en golang con el paquete reflect
	var numero int = 3
	var nombre string = "Jordy"
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(nombre))
}
