package main

import "fmt"

// func main(){
// 	ii := []int{1,2,3,4,5,6,7,8}
// 	n:= foo(ii...)
// 	defer fmt.Println(n)

// 	io := []int{1,2,3,4,5,6,7,8,9}
// 	s:= bar(io)
// 	fmt.Println(s)
// }
// func foo(vp ...int)int{
// 	total := 0
// 	for _, v := range vp{
// 		total += v 
// 	}
// 	return total
// }
// func bar (vp[]int)int{
// 	total := 0
// 	for _, v := range vp{
// 		total += v 
// 	}
// 	return total
// }
func main(){
	defer foo()
	bar()
	mar()
}
func foo(){
	fmt.Println("Print foo")
}
func bar(){
	fmt.Println("Print bar")
}
func mar(){
	fmt.Println("Print mar")
}