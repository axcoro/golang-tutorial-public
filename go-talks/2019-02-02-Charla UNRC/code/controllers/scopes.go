package controllers

import "github.com/mercadolibre/...../src/api/services"

func main() {

	services.SomeVar = 5  // OK
	services.otherVar = 3 // compile error
}
