package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDte string
}

// Algunas funciones pueden funcionar como metodos.
// Go lo hace de manera implicita a diferencia de otros lenguajes.
func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporatyEmployee struct {
	Person
	Employee
	taxRate int
}

// Se implementa la interfaz.
func (tEmployee TemporatyEmployee) getMessage() string {
	return "Temporaty Employee"
}

// Es una interface:
type PrintInfo interface {
	getMessage() string
}

// Se crea el metodo de la interfaz
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporatyEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
