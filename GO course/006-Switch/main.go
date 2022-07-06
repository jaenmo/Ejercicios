package main

import (
	"fmt"
)

func main() {
	day := "viernes"
	switch day {
	case "lunes", "martes", "miernes":
		fmt.Printf("%s es un buen dia", day)
	default:
		fmt.Printf("%s es otro dia", day)
	}
}
