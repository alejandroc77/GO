package main

import "fmt"

func main() {
	myslice := []string{"bolvo", "mazda", "toyota"}
	fmt.Println(len(myslice)) //longitud del slice
	fmt.Println(cap(myslice)) //capacidad del slice
	fmt.Println(myslice[1])
}
