package main

import "fmt"

func main(){
  defer foo()
  boo()
  moo()
// 	x := sum (1, 2, 3, 4, 5, 6, 7, 8, 9)
// 	fmt.Println("the total is", x)



// }
// func sum (x ...int)int{
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)

// 	sum := 0
// 	for i, v := range x {
// 		sum += v
// 		fmt.Println("for the number in position", i, "we are adding", v, "which is", sum)
// 		}
// return sum
	}
	func foo(){
		fmt.Println("foo")
	}
	func boo(){
		fmt.Println("boo")
	}
	func moo(){
		fmt.Println("moo")
	}