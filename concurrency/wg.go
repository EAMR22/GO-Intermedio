package main

//	import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {

// 	var wg sync.WaitGroup // Esto es un contador.

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)              // Cada vez que enviamos una go routine, le sumamos 1 alcontador.
// 		go doSomething(i, &wg) // Se tiene que enviar por referencia, porque si no se pasa como copia.
// 	}

// 	wg.Wait() // Bloquea el programa, hasta que el contador llegue a 0.
// }

// func doSomething(i int, wg *sync.WaitGroup) {
// 	defer wg.Done() // Cada vez que termina le restamos 1 al contador.
// 	fmt.Printf("Started %d\n", i)
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("End")
// }//
