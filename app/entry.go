package main

import (
	"fmt"
	"log"
	"math/rand"

	"entry/greet"
)

func main() {

	err := somehting()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greet.Morning)
}

func somehting() (_ error) {
	return
}

type int string

func (int) MarshalJSON() ([]byte, error) {
	if rand.Intn(10000) == 9999 {
		panic("have fun debuggin this, fucker!")
	}
	return nil, nil
}
