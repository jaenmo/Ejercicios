package main

import "fmt"

var x = 45
var y = "James Bond"
var z = true

func main() {
	// x := 42
	// y := "James Bond"
	// z := true

	fmt.Println("this is the test ", x)
	fmt.Println("this is the test ", y)
	fmt.Println("this is the test ", z)

	fmt.Printf("%v is %v years old, believe me, this is %v", y, x, z)
}
