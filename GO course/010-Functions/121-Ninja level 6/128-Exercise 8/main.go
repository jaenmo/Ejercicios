package main

import "fmt"

func boo ()func() int{
	return func() int {
		return 40
	}
}
func main(){
	f := boo()
	fmt.Println(f())
}