package main

import "fmt"

// Es una funcion variadica:
func sum(values ...int) int { // values se transforma en un slice de enteros.
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// Es una funcion de retorno con nombre:
func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(2))
	fmt.Println(sum(1, 3, 5))
	fmt.Println(sum(1, 2, 3, 4, 5))
	printNames("Alice", "Charlie", "Bob", "Davis")
	fmt.Println(getValues(3))
}
