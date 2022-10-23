package main

import "fmt"

func main() {
	myslice := []string{"bolvo", "mazda", "toyota"}
	fmt.Println(len(myslice)) //longitud del slice
	fmt.Println(cap(myslice)) //capacidad del slice
	fmt.Println(myslice[1])
	anexar()
	anexar2()
}

func anexar() {
	myanexo := []int{1, 2, 3, 4, 5}
	fmt.Println(myanexo)

	myanexo = append(myanexo, 12, 3)
	fmt.Println(myanexo)
}

// otra forma de anexar

func anexar2() {
	fmt.Println("==== anexar 2 ======")
	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)
	fmt.Println(myslice3)

}
