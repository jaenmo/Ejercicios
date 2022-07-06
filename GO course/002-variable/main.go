package main

import "fmt"

// var language = "Go"

/* var (
	language = "Go"
	company  = "Google"
	year     = 2009
)*/

func main() {

	for i := 50 ; i < 250 ; i++ {
		fmt.Printf("%T\t -%d\t- %b\t- %x\t- %X\t - %q\n", i, i, i, i, i, i)
	}

	// fmt.Println("Aprendiendo variables en " + language)
	// fmt.Printf("Aprendiendo variables en %s, un lenguaje creado por %s en el year %d ", language, company, year)

	language := "Go"
	fmt.Println("Aprendiendo variables en " + language)
}
