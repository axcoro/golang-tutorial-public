package main

import (
	"fmt"
)

func main() {
	i := 42

	p := &i         // puntero a i
	fmt.Println("valor de i: ", *p)
	*p = 21         // setea i a través del puntero
	fmt.Println("nuevo valor de i: ", i)

	p = new(int)   // puntero a un valor de tipo int
	*p = 10
	fmt.Println("direccion de memoria de p: ", p)
	fmt.Println("valor de p: ", *p)
	
	//p++ //error! a diferencia de C, Go no acepta aritmetica de punteros
}