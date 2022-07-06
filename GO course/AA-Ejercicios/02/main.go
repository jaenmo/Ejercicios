package main

import "fmt"

func main(){
	a := []int{3, 35, 21, 66, 87}
	max := a[1]
	min := a[2]

	for _, v := range a{
		if max < v{
			max = v
		}
		if min > v{
			min = v
		}
	}
	fmt.Println(max, min)
	
	fmt.Println("------------")
	
	fmt.Println(m(2, 65, 87, 3))
	greatest := m(34, 5, 231, 25, 543, 65)
	fmt.Println(greatest)

}

func m (num ...int) int{
	var largest int
	for _, v := range num{
		if v > largest{
			largest = v
		}
	}
	return largest
}