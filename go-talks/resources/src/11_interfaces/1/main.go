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

func main() {
	var superman SuperHeroe = Superman{Alias: "Superman", Edad: 28}

	superman.Poderes()
}
