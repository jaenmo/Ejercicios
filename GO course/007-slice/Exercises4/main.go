package main

import "fmt"

func main() {
	// // ### EXERCISE 1
	// array1 := [5]int{}
	// //fmt.Println(array1)
	// array1 = [5]int{1, 2, 3, 5, 8}
	// //fmt.Println(array1)

	// for i, value := range array1 {
	// 	fmt.Println(i, value)
	// }
	// fmt.Printf("%T", array1)

	// ### EXERCISE 2

	// slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// fmt.Println(slice)
	
	// for i, value := range slice {
	// 	fmt.Println(i, value)
	// }
	// fmt.Printf("%T", slice)

	// ### EXERCISE 3

	// slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// slice1 := append(slice[:5])
	// fmt.Println(slice1)
	// slice2 := append(slice[5:])
	// fmt.Println(slice2)
	// slice3 := append(slice[2:7])
	// fmt.Println(slice3)
	// slice4 := append(slice[1:6])
	// fmt.Println(slice4)

	// fmt.Println(slice[:5])
	// fmt.Println(slice[5:])
	// fmt.Println(slice[2:7])
	// fmt.Println(slice[1:6])

	// ### EXERCISE 4

	// slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// slice = append(slice, 52)
	// fmt.Println(slice)
	// slice = append(slice, 53, 54, 55)
	// fmt.Println(slice)
	// slice5 := []int{56, 57, 58, 59, 60}
	// fmt.Println(slice5)
	// slice = append(slice, slice5...)
	// fmt.Println(slice)
	// slice = append(slice5, slice[2])
	// fmt.Println(slice)

	// ### EXERCISE 5
	// slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// slice = append(slice[:3], slice[6:]...)
	// fmt.Println(slice)

	// ### EXERCISE 6
	// slice := []string{
	// 	`Alabama (AL)`,
	// 	`Alaska (AK)`,
	// 	`Arizona (AZ)`,
	// 	`Arkansas (AR)`,
	// 	`California (CA)`,
	// 	`Colorado (CO)`,
	// 	`Connecticut (CT)`,
	// 	`Delaware (DE)`,
	// 	`Florida (FL)`,
	// 	`Georgia (GA)`,
	// 	`Hawaii (HI)`,
	// 	`Idaho (ID)`,
	// 	`Illinois (IL)`,
	// 	`Indiana (IN)`,
	// 	`Iowa (IA)`,
	// 	`Kansas (KS)`,
	// 	`Kentucky (KY)`,
	// 	`Louisiana (LA)`,
	// 	`Maine (ME)`,
	// 	`Maryland (MD)`,
	// 	`Massachusetts (MA)`,
	// 	`Michigan (MI)`,
	// 	`Minnesota (MN)`,
	// 	`Mississippi (MS)`,
	// 	`Missouri (MO)`,
	// 	`Montana (MT)`,
	// 	`Nebraska (NE)`,
	// 	`Nevada (NV)`,
	// 	`New Hampshire (NH)`,
	// 	`New Jersey (NJ)`,
	// 	`New Mexico (NM)`,
	// 	`New York (NY)`,
	// 	`North Carolina (NC)`,
	// 	`North Dakota (ND)`,
	// 	`Ohio (OH)`,
	// 	`Oklahoma (OK)`,
	// 	`Oregon (OR)`,
	// 	`Pennsylvania (PA)`,
	// 	`Rhode Island (RI)`,
	// 	`South Carolina (SC)`,
	// 	`South Dakota (SD)`,
	// 	`Tennessee (TN)`,
	// 	`Texas (TX)`,
	// 	`Utah (UT)`,
	// 	`Vermont (VT)`,
	// 	`Virginia (VA)`,
	// 	`Washington (WA)`,
	// 	`West Virginia (WV)`,
	// 	`Wisconsin (WI)`,
	// 	`Wyoming (WY)`}

	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	// for i, value := range slice {
	// 	fmt.Printf("%v\t%v\n", i, value)
	// }

	// ### EXERCISE 7
	// slice1 := []string{"James", "Bond", "Shaken not stirred"}
	// slice2 := []string{"Miss", "Moneypenny", "Hellooooooo, James"}
	// s := [][]string{slice1, slice2}
	// fmt.Println(s)

	// for i, value1 := range s {
	// 	fmt.Println("Record:", i)
	// 	for j, value2 := range value1 {
	// 		fmt.Printf("\t index position: %v \t value:\t %v\n", j, value2)
	// 	}
	// }

	// ### EXERCISE 8

	// slice1 := []string{"Shaken, not stirred", "Martinis", "Women"}
	// slice2 := []string{"James Bond", "Literature", "Computer Science"}
	// slice3 := []string{"Being evil", "Ice cream", "Sunsets"}
	// m := map[string][]string{
	// 	"James Bond":      slice1,
	// 	"Moneypenny_miss": slice2,
	// 	"No_dr":           slice3,
	// }
	// fmt.Println(m)
	// for key, value := range m {
	// 	fmt.Println("For the record of", key)
	// 	for i, value2 := range value {
	// 		fmt.Println("\t", i, value2)
	// 	}
	// }
	// fmt.Println("#########")

	// // ### EXERCISE 9
	// slice4 := []string{"matar", "Asel el amol"}
	// m["villano"] = slice4

	// for key1, value3 := range m {
	// 	fmt.Println("Records", key1)
	// 	for key2, value4 := range value3 {
	// 		fmt.Println("\t", key2, value4)
	// 	}
	// }
	// fmt.Println("#########")

	// // ### EXERCISE 10
	// delete(m, "James Bond")
	// for key1, value3 := range m {
	// 	fmt.Println("Records", key1)
	// 	for key2, value4 := range value3 {
	// 		fmt.Println("\t", key2, value4)
	// 	}
	// }

	// ### EXERCISE 11 (extra)

	// slice1 := []string{"uno", "dos", "tres", "cuatro"}
	// slice2 := []string{"1", "2", "3", "4"}
	// slice3 := []string{"1", "2", "3", "4"}
	// m1 := map[float64][]string{
	// 	3.14: slice1,
	// 	5.25: slice2,
	// 	7.46: slice3,
	// 	1.10: []string{},
	// }
	// for k1, v1 := range m1 {
	// 	fmt.Println("Dato inicial", k1)
	// 	for _, v2 := range v1 {
	// 		fmt.Println("\t", v2)
	// 	}
	// }

	// size := 0
	// fmt.Print("Number of elements n=")
	// fmt.Scanln(&size)
	// fmt.Println("Enter array elements")
	//

}
