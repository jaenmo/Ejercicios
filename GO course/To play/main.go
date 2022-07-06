package main

import (
	"fmt"
)

func main() {

	// x := []int{18, 31, 52, 83}
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(x[0])

	// for i := range x {
	// 	fmt.Println(i)
	// }
	semana := []string{"lunes", "martes", "miercoles"}
	fmt.Println(semana)
	semana = append(semana, "jueves", "viernes")
	fmt.Println(semana)

	y := []string{"sabado", "domingo"}
	fmt.Println(semana, y)
	semana = append(semana, y...)
	fmt.Println(semana, y)

}
