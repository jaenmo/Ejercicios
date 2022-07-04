package main

import "fmt"

var l float64
var a float64
var g float64

const (dWood = 1.5
 	dMarble = 2.52
 	dPlastic = 0.9)

func main(){

mesa := map [string]float64{
	"length": l,
	"width": a,
	"thickness": g,
	"Wood": dWood,
	"Marble": dMarble,
	"Plastic": dPlastic,
	}
fmt.Println(mesa)

introduceData()
chooseMaterial()
calculateWeight()

}

 func introduceData(){
	fmt.Println("Introduce length cm")
	fmt.Scanln(l)
	
	fmt.Println("Introduce width in cm")
	fmt.Scanln(a)
	
	fmt.Println("Introduce thickness in cm")
	fmt.Scanln(g)
 }

 func chooseMaterial(){
	fmt.Println("which material is it?")
	material := map [int]string{
		1: "wood",
		2: "marble",
		3: "plastic",
	}
	fmt.Println(material)
 }

 func calculateWeight(){
	 wWood := l * a * g * dWood
	 wMarble := l * a * g * dMarble
	 wPlastic := l * a * g * dPlastic

	 fmt.Println("Weight of table in wood is", wWood)
	 fmt.Println("Weight of table in marble is", wMarble)
	 fmt.Println("Weight of table in plastic is", wPlastic)


 }
