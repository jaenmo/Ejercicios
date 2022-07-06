package main

import "fmt"

func main() {

	//// defined struct

	// type person struct {
	// 	first string
	// 	last  string
	// 	year  int
	// }
	// es := person{
	// 	first: "JD",
	// 	last:  "Wu",
	// 	year:  18,
	// }
	// fmt.Println(es)

	//// not defined struct

	es := struct {
		first string
		last  string
		year  int
	}{
		first: "JD",
		last:  "Wu",
		year:  18,
	}
	fmt.Println(es)
}
