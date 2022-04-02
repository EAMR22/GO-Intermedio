package main

import "fmt"

func Generator(c chan<- int) { // Cuando ponemos la <- a la derecha del canal, nos dice que es de escritura.
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

func Print(c <-chan int) { // Cuando indicamos la <- al lado izquierdo del canal, significa que es de lectura.
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles) // Lo utilizamos sin la palabra reservada go para que el programa se bloquee.
}
