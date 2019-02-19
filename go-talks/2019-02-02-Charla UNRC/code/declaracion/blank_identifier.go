package main

import "fmt"

func main() {
	numeros := []string{"cero", "uno", "dos", "tres"}
	for _, valor := range numeros {
		fmt.Println(valor)
	}
}
