package main

import "fmt"

func main() {
	v := 40
	fmt.Println(v)
	fmt.Println(&v)

	var b *int = &v
	fmt.Println(b)
	fmt.Println(*b)

	*b = 45
	fmt.Println(v)
	fmt.Println(&v)

}
