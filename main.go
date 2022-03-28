package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int // Lo hace de manera explicita.
	x = 8
	y := 7 // Aqui de manera implicita.

	fmt.Println(x)
	fmt.Println(y)

	// Errores:
	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil { // Maneja el error de manera explicita.
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// map:
	m := make(map[string]int) // Mapea llaves de tipo string a valores de tipo entero.
	m["key"] = 6
	fmt.Println(m["key"])

	// slice:
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	// canal:
	c := make(chan int)
	go doSomething(c)
	<-c

	//apuntadores:
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)  // Imprime la direccion de memoria donde esta alojado "g".
	fmt.Println(*h) // Accede al valor que tiene "g".
}

func doSomething(c chan int) { // Esta funcion va a correr en una rutina diferente a la de main.
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1 // Enviara el valor de 1.
}
