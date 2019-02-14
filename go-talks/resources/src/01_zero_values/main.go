package main

import "fmt"

var i int
var f float64
var b bool
var s string
var arr [1]float32

func init() { fmt.Println("Se ejecuta init") }

func main() { fmt.Printf("%v %v %v %q %v\n", i, f, b, s, arr) }
