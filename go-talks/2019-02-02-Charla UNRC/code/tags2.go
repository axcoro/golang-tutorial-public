package main

import (
	"fmt"
	"reflect"
)

type Persona struct {
	Nombre string `custom:"algo"`
}

func main() {

	p := Persona{Nombre: "Test"}

	field, _ := reflect.TypeOf(p).FieldByName("Nombre")
	tag := field.Tag.Get("custom")

	fmt.Printf("el valor del tag es: '%v'", tag)
}
