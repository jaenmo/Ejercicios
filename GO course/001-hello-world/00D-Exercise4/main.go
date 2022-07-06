package main

import (
	"fmt"
)

type mostro int

var x mostro

func main() {
	x = 42
	fmt.Printf("wich type is %T\n", x)
	fmt.Println(x)

}
