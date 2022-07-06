package main

import "fmt"

func main() {

	x := "hola que tal"
	xs := []byte(x)
	fmt.Println(x)
	fmt.Println(xs)
	fmt.Printf("%#U\n", xs)

	x = "a gusto"
	fmt.Println(x)
}
