package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select { // select nos permite tener diferentes casos, cuando estamos trabajando con diferentes canales,
		case channelMsg1 := <-c1: // y necesitamos saber cual es el canal que a sido activado.
			fmt.Println(channelMsg1)
		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param // Envia param a traves del canal.
}
