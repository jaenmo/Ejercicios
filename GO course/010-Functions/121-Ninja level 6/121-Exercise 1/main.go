package main

import (
	"fmt"
)

func main(){

// OPTION 1

// 	foo()
// 	bar()

// }
// func foo (){
// 	x := 40
// 	fmt.Println(x)

// }
// func bar (){
// 	x := 5
// 	y := "test"
// 	fmt.Println(x)
// 	fmt.Println(y)

// OPTION 2

n := foo()
x, s := bar()
fmt.Println(n, x, s)

}
func foo ()int{
	return 42
}
func bar ()(int,string){
	return 1984, "Big Brother"
}