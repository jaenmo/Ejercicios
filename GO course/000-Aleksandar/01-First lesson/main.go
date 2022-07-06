package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := []string{"1", "2", "3f"}

	var ints []int
	for i, s := range str {
		fmt.Printf("index: %d\n", i)

		num, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			fmt.Println(err)
			continue
		}

		ints = append(ints, int(num))
	}

	fmt.Println(ints)
}
