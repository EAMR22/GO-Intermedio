package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// Es una composicion sobre la herencia:

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
}
