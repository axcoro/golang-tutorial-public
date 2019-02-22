package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}
type Student struct {
	Human      // El struct Student va a incluir los campos que tiene Human.
	speciality string
}

func main() {
	// inicializamos a student
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	// campos accesibles
	fmt.Println("Su nombre es ", mark.name)
}
