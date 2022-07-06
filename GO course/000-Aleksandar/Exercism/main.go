package main

import "fmt"

func main(){
	a := myNameIs()

	fmt.Println(a)
}

func myNameIs (){
	fmt.Println("Hello, my name is")
}