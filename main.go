package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int // Lo hace de manera explicita.
	x = 8
	y := 7 // Aqui de manera implicita.

	fmt.Println(x)
	fmt.Println(y)

	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil { // Maneja el error de manera explicita.
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	m := make(map[string]int) // Mapea llaves de tipo string a valores de tipo entero.
	m["key"] = 6
	fmt.Println(m["key"])

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
}
