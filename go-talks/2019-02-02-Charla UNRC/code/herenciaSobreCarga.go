    package main
    import "fmt"
    type Human struct {
        name string
         age int
    }
    type Student struct {
        Human
        school string
    }
    type Employee struct {
        Human
    }
    func (h *Human) SayHi() {
        fmt.Printf("Hola, Yo soy %s", h.name)
    }
    func (e *Student) SayHi() {
        fmt.Printf("Hola, estudio en %s", e.school)
    }
    func main() {
        mark := Student{Human{"Mark", 25}, "MIT"}
        sam := Employee{Human{"Sam", 45}}
        mark.SayHi()
        sam.SayHi()
    }
