package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo desde Go")

	foo()

	fmt.Println("y esto va al final")
}
func foo() {
	fmt.Println("esto va a la mitad")

	n, err := fmt.Println("Hello, playmate, 18, false")
	fmt.Println(n)
	fmt.Println(err)
}
