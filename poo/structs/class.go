package main

import "fmt"

// Estructura de una clase en go:

type Employe struct {
	id   int
	name string
}

// Creacion de metodos en go:

func (e *Employe) SetId(id int) { // El struct Employe va a poseer un metodo llamado SetId.
	e.id = id
}

func (e *Employe) SetName(name string) {
	e.name = name
}

func (e *Employe) GetId() int {
	return e.id
}

func (e *Employe) GetName() string {
	return e.name
}

func main() {
	e := Employe{}
	// fmt.Printf("%v", e)
	e.id = 1
	e.name = "name"
	// fmt.Printf("%v", e)
	e.SetId(5)
	e.SetName("goolang")
	// fmt.Printf("%v", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
