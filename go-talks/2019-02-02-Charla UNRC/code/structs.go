package main

import "fmt"

type Persona struct {
	ID       int
	Nombre   string
	Apellido string
	Email    string
}

func main() {

	p := Persona{Nombre: "Juan", Apellido: "Perez"}

	fmt.Println("Nombre:", p.Nombre)
	fmt.Println("Apellido:", p.Apellido)
}
