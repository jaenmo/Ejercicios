package main

import (
	"fmt"
	"math"
)

type circle struct{
	radius float64
}
type rectangle struct{
	length 	float64
	width	float64
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}
func (s rectangle) area() float64{
	return s.length * s.width
}

type shape interface {
	area () float64
}
func info (s shape) {
	x := s.area()
	fmt.Println(x)
}
func main(){
	c := circle {
		radius: 12.345,
	}
	s := rectangle {
		length: 15.1,
		width: 5.2,
	}

	info(c)
	info(s)

}
