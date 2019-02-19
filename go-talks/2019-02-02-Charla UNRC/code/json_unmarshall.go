package main

import "encoding/json"
import "fmt"

type Persona struct {
	Nombre   string `json:"name"`
	Apellido string
}

func main() {
	jsonStr := []byte(`{"id":123456,"tags":["a","b"]}`)
	var jsonMap map[string]interface{}

	if err := json.Unmarshal(jsonStr, &jsonMap); err != nil {
		panic(err)
	}
	fmt.Println(jsonMap)

	id := jsonMap["id"].(float64)
	fmt.Println(id)

	tags := jsonMap["tags"].([]interface{})
	tag := tags[0].(string)
	fmt.Println(tag)

	str := `{"name": "A", "Apellido": "B"}`
	persona := Persona{}
	json.Unmarshal([]byte(str), &persona)
	fmt.Println(persona)
	fmt.Println(persona.Nombre)
}
