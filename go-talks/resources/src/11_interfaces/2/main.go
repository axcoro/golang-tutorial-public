package main

import (
	"fmt"
)

type SuperHero interface {
	Poderes()
}

type Superman struct {
	Alias string
	Age   int
}

func (s Superman) Poderes() {
	fmt.Printf("Poderes de %s: %s", s.Alias, "invulnerabilidad, vision de rayos x, volar, velocidad")
}

type Batman struct {
	Alias string
	Age   int
}

func (s Batman) Poderes() {
	fmt.Printf("Poderes de %s: %s", s.Alias, "dinero")
}

func main() {
	var superman SuperHero = Superman{Alias: "Superman", Age: 28}

	superman.Poderes()

	fmt.Println("")

	var batman SuperHero = Batman{Alias: "Batman", Age: 24}

	batman.Poderes()

	fmt.Println("")
}
