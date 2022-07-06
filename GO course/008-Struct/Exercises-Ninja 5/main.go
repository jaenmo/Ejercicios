package main

import (
	"fmt")


func main() {


	



	// for a := 1; a <= 100; a++ {

	// 	if a%3 == 0 {
	// fmt.Println("Fizz")
	// 	}
	// 	if a%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	}
	// 	if a%5 == 0 && a%3 == 0{
	// 		fmt.Println("FizzBuzz")
	// 	}
	// 	if a%5 != 0 && a%3 != 0 {
	// 		fmt.Println(a)

	// 	}

	// -------EXTRA EXERCISE------
// var sum int

// for b := 1; b < 10; b++ {
// 	if b%3 == 0 || b%5 == 0 {
// 		//sum += b
// 		fmt.Println(b)
// 	}

// }
// fmt.Println(sum)



// x := 0
// 	for x <= 10 {
// 		x++

// 		if x > 10{
// 			break
// 		}
// 		if x%2 == 0 {
// 			continue
// 		} 
// 		fmt.Println(x)

// 	}
// 	fmt.Println("#######")

// 	for y := 0; y<=10; y++{
// 		if y%2 == 0 {
// 		fmt.Println(y)
// 	}
// 	}
		
		
		


	// ### EXERCISE 1

	// type person struct {
	// 	firstName string
	// 	lastName	string
	// 	favIceCreamFlav	[]string
	// }
	// p1 := person {"James", "Bond", []string{"Vodka", "Martini", "Heineken"}}
	// p2 := person {"Miss", "Moneypenny", []string {"Melon", "Mango", "Lima"}}
	
	// fmt.Println(p1)
	// fmt.Println(p2)
	
	// for i, value := range p1.favIceCreamFlav{
	// 	fmt.Println(i, value)
	// } 

	// type person struct {
	// 	first_name                  string
	// 	last_name                   string
	// 	favourite_ice_cream_flavour []string
	// }

	// p1 := person{
	// 	first_name:                  "James",
	// 	last_name:                   "Bond",
	// 	favourite_ice_cream_flavour: []string{"vanila", "strawberry"},
	// }
	// var p2 person
	// p2.first_name = "Miss"
	// p2.last_name = "Moneypenny"
	// p2.favourite_ice_cream_flavour = []string{"chocolate", "melon"}
	// p2 := person{
	// 	first_name:                  "Miss",
	// 	last_name:                   "Moneypenny",
	// 	favourite_ice_cream_flavour: []string{"chocolate", "melon"},
	// }
	// fmt.Println(p1.first_name)
	// fmt.Println(p2.first_name)
	// fmt.Println(p2.favourite_ice_cream_flavour)
	// fmt.Println(p1, p2)

	// fmt.Println(p1.first_name, p1.last_name)

	// for _, value := range p1.favourite_ice_cream_flavour {
	// 	fmt.Println("\t", value)
	// }

	// fmt.Println(p2.first_name, p2.last_name)

	// for _, value1 := range p2.favourite_ice_cream_flavour {
	// 	fmt.Println("\t", value1)
	// }

	// // ### EXERCISE 2

	// m1 := map[string]person{
	// 	p1.last_name: p1,
	// 	p2.last_name: p2,
	// }
	// // fmt.Println(m1)

	// for key1, value := range m1 {
	// 	//fmt.Println(key1, value)
	// 	fmt.Println(key1)
	// 	fmt.Println(value)

	// 	fmt.Println("---------")

	// }

	// type individual struct{
	// 	nombre string
	// 	apellidos string
	// 	edad int

	// type group struct{
	// 	individual
	// 	location string
	// }

	// var a group

	// a := individual


		// // ### EXERCISE 3

//  type vehicle struct{
// 	 doors int
// 	 color string
//  }
//  type truck struct{
// 	vehicle
// 	fourWheel bool
//  }
//  type sedan struct{
// 	 vehicle
// 	 luxury bool
//  }

//  t := truck{
// 	vehicle: vehicle{
// 		doors: 2,
// 		color: "black",
// 	},
// 	fourWheel: true,
//  } 

//  s := sedan{
// 	vehicle: vehicle{
// 		doors: 4,
// 		color: "blue",
// 	},
// 	luxury: true,
//  } 
//  fmt.Println(t)
//  fmt.Println(s)

//  fmt.Println(t.doors, t.color, t.fourWheel)
//  fmt.Println(s.doors, s.color)


		// // ### EXERCISE 4

	ex := struct{
		first string
		friends map[string]int
		favDrink []string
	}{
		first: "james",
		friends: map[string]int{
			"Rose": 2,
			"Rachel": 4,
			"Joey": 8,
		},
		favDrink: [] string {
			"Vodka",
			"cola",
			"Beer",
		},
	}
fmt.Println(ex)
fmt.Println(ex.first)
fmt.Println(ex.friends)
fmt.Println(ex.favDrink)

for k, v := range ex.friends{
	fmt.Println(k,v)
}
for k1, v1 := range ex.favDrink{
	fmt.Println(k1,v1)
}



}