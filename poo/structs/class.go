package main

import "fmt"

type Employe struct { // struct es equivalente a una clase en otros lenguajes de programacion.
	id   int
	name string
}

func main() {
	e := Employe{} // Aqui instanciamos el objeto.
	fmt.Printf("%v", e)
	e.id = 1
	e.name = "name"
	fmt.Printf("%v", e)
}
