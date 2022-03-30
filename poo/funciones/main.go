package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int { // Es una funcion anonima, a la que solo se utilizara una vez.
		return x * 2
	}() // El "()" llama a la funcion.
	fmt.Println(y)

	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c
}
