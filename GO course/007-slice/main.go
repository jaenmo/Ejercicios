package main

import (
	"fmt"
)

func main() {
// create a slice and options

	// x := []int{18, 31, 52, 83}
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(x[0])

// range and index

	// for i := range x {
	// 	fmt.Println(i)
	// }

// create a slice with strings
// add elements to the slice

	semana := []string{"lunes", "martes", "miercoles"}
	fmt.Println(semana)
	semana = append(semana, "jueves", "viernes")
	fmt.Println(semana)

//add another slice to the slice

	y := []string{"sabado", "domingo"}
	fmt.Println(semana, y)
	semana = append(semana, y...)
	fmt.Println(semana)

//remove elements from the slice

//just show elements in the middle
	// semana = append(semana[2:4])
	// fmt.Println(semana)
	// // fmt.Println(len(semana))

// remove elements in the middle
	semana = append(semana[:2], semana[4:]...)
	fmt.Println(semana)

// make the array bigger
	adicional := make([]string, 10, 11)
	adicional[0] = "lunes"
	adicional[1] = "martes"
	adicional[9] = "domingo"

	fmt.Println(adicional)
	fmt.Println(len(adicional))
	fmt.Println(cap(adicional))

	adicional = append(adicional, "enero", "febrero")
	fmt.Println(adicional)
	fmt.Println(len(adicional))
	fmt.Println(cap(adicional))

	adicional = append(adicional, "marzo")
	adicional = append(adicional, "abril")
	fmt.Println(adicional)
	fmt.Println(len(adicional))
	fmt.Println(cap(adicional))

//slice multidimensional
// ####### check how to combine strings with integers in a multidimensional slice

	// s1 := []string{"Julio", "Agosto", "Septiembre"}
	// s2 := []int{7, 8, 9}
	// s3 := []int{0, 1, 2}
	// fmt.Println(s1, s2)
	// s4 := [][]int{s2, s3}
	// fmt.Println(s4)
	// ## WRONG## s4a := [string][]int{s1, s2}

// maps
	fmt.Println("####################")
	m := map[string]int{
		"Enero":   01,
		"Febrero": 02,
		"Marzo":   03,
		"Abril":   04,
	}

	fmt.Println(m)
	fmt.Println(m["Enero"])
	fmt.Println(m["Febrero"])
	fmt.Println(m["Mayo"])
fmt.Println("------------------")
	nv, ok := m["Mayo"]
	fmt.Println(nv)
	fmt.Println(ok)
	
fmt.Println("------------------")

	if _, ok := m["julio"]; ok {
		fmt.Println("If this is ok, print", ok)
	}
	m["Junio"] = 06
	for v2, nv := range m {
		fmt.Println(v2, nv)
	}
	delete(m, "Marzo")
	fmt.Println(m)

}
