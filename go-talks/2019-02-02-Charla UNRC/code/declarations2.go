package main

import (
	"fmt"
)

type Currency string

func main() {

	c := Currency("ARS")                                 // tipo(valor)
	fmt.Printf("currency: %v, string: %v", c, string(c)) // casteo al tipo base

}
