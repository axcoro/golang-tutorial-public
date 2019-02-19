package main

import "fmt"

type Animal interface {
	hacerRuido()
}

type Pato struct{}

func (p Pato) hacerRuido() {
	fmt.Printf("cuack cuack!")
}

func main() {
	var a Animal = Pato{}
	a.hacerRuido()
}
