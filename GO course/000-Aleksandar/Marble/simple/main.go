package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter the width (cm): ")
	var width float64
	fmt.Scanln(&width)

	fmt.Println("Enter the lengh (cm): ")
	var length float64
	fmt.Scanln(&length)

	fmt.Println("Enter the thickness (cm): ")
	var thickness float64
	fmt.Scanln(&thickness)

	fmt.Print("The weight of the board is:\n")

	// density of marble 2.7 g/cm3
	density := 2.7
	fmt.Print(width * length * thickness * density, "g\n")
	fmt.Print(width * length * thickness * density / 1000, "kg")
}
