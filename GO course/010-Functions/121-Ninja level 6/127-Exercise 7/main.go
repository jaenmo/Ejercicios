package main

import "fmt"

func main(){

x := func (){
	i := "Hola que tal"
	fmt.Println(i)
	}
	x()
	fmt.Printf("%T\n", x)
}