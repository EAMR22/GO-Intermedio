package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{ // La función regresa con & para que regrese la referencia y no una copia, de la struct.
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// forma 1:
	e := Employee{}
	fmt.Printf("%v\n", e)

	// forma 2:
	e2 := Employee{
		id:       1,
		name:     "name",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	// forma 3:
	e3 := new(Employee) // Cuando se usa new devuelve un apuntador de la instancia Employee.
	fmt.Printf("%v\n", *e3)
	e3.id = 1
	e3.name = "name"
	fmt.Printf("%v\n", *e3)

	// forma 4:
	e4 := NewEmployee(1, "name", true)
	fmt.Printf("%v\n", *e4)
}
