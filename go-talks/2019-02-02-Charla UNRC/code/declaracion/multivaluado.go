package main

import "fmt"

func main() {
	a, b := palabras()
	fmt.Println(a, b)
	c, _ := palabras()
	fmt.Println(c)
}

func palabras() (string, error) {
	return "valor de retorno", fmt.Errorf("Error")
}
