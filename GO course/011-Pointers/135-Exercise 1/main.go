package main

import "fmt"

func zero(x *int) { // why we have an address of x when we didn't asign any value yet?
	// x = 0
	fmt.Println(x)
	//fmt.Println(*&x)
	*x = 0 //if we have the x = 0 on top, the results are different

}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)
}
