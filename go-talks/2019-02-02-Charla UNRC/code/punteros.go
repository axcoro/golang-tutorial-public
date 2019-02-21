package main

import "fmt"

func main() {
	edad := 25
	var puntero *int //Declaracion de un puntero

	puntero = &edad // puntero a edad
	fmt.Println("edad: ", *puntero)
	*puntero = 21 // setea edad a través del puntero
	fmt.Println("nuevo valor de edad: ", edad)
	fmt.Println("direccion de memoria de puntero: ", puntero)

	//puntero++ //error! a diferencia de C, Go no acepta aritmetica de punteros
}
