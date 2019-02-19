package main

import "encoding/json"
import "fmt"

type Persona struct {
	Nombre   string `json:"name"`
	Apellido string
}

func main() {
	m1, _ := json.Marshal(true)
	fmt.Println(string(m1))

	m2, _ := json.Marshal(1)
	fmt.Println(string(m2))

	m3, _ := json.Marshal([]string{"a", "b", "c"})
	fmt.Println(string(m3))

	m4, _ := json.Marshal(map[string]int{"manzana": 5, "banana": 7})
	fmt.Println(string(m4))

	m5, _ := json.Marshal(&Persona{Nombre: "Nombre", Apellido: "Apellido"})
	fmt.Println(string(m5))
}
