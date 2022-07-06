package main

import "fmt"

var l float64
var a float64
var g float64

const (dWood = 1.5
 	dMarble = 2.52
 	dPlastic = 0.9)

func main(){

fmt.Println("this program is to calculate the weight of a table, please follow the steps")

introduceData()
chooseMaterial()

}

 func introduceData(){
	fmt.Println("Introduce length cm")
	fmt.Scanln(&l)
	
	fmt.Println("Introduce width in cm")
	fmt.Scanln(&a)
	
	fmt.Println("Introduce thickness in cm")
	fmt.Scanln(&g)
 }

 func chooseMaterial(){
	wWood := l * a * g * dWood
	wMarble := l * a * g * dMarble
	wPlastic := l * a * g * dPlastic

	fmt.Println("which material is it?")
	material := []string{"wood", "marble", "plastic"}
	for i, v := range material{
		fmt.Println(i+1, v)
	}

	var i int
	fmt.Scanln(&i)
	
    switch i {
    case 1:
        fmt.Println("You have chosen wood")
		fmt.Printf("Weight of table in wood is %v g\n", wWood)
		fmt.Println(wWood / 1000, "Kg")

    case 2:
        fmt.Println("You have chosen marble")
		fmt.Printf("Weight of table in marble is %v g\n", wMarble)
		fmt.Println(wMarble / 1000, "Kg")
    case 3:
        fmt.Println("You have chosen plastic")
		fmt.Printf("Weight of table in plastic %v g/n", wPlastic)
		fmt.Println(wPlastic / 1000, "Kg")
	}
	
 }

