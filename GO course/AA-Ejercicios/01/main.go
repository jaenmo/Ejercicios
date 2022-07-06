package main

import "fmt"

// func half (x int)(y int, z bool){
// 	y = x/2
// 	if y%2 == 0 {
// 		fmt.Println(y, false)
// 	}else{
// 		fmt.Println(y, true)
// 	}
// 	return 
// }

// func main(){
// 	half(1)
// 	half(2)
// }

// func half(n int)(float64, bool, float64, float64){
// 	return float64(n) / 2, n%2 == 0, float64(n)*2, float64(n)*float64(n)
// }

// func main(){
// 	h, even, z, d := half (5)
// 	fmt.Println(h, even, z, d)
// }



func main(){
	half := func(n int)(float64, bool, float64, float64){
		return float64(n) / 2, n%2 == 0, float64(n)*2, float64(n)*float64(n)
	}
	// h, even, z, d := half (5)
	fmt.Println(half (5))
}
