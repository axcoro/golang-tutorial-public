package main

import (
	"fmt"
)

type SuperHeroe interface {
	Poderes()
}

type Superman struct {
	Alias string
	Edad  int
}

func (s Superman) Poderes() {
	fmt.Printf("Poderes de %s: %s", s.Alias, "invulnerabilidad, vision de rayos x, volar, velocidad")
}

type SuperBoy struct {
	Superman
}

func (s *SuperBoy) Poderes() {
	fmt.Printf("Poderes de %s: %s", s.Alias, "volar")
}

func main() {
	var superboy SuperBoy = SuperBoy{}

	superboy.Alias = "Superboy"
	superboy.Edad = 24

	superboy.Poderes()

	fmt.Println("")
}
