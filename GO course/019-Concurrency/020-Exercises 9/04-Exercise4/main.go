package main

import (
	"fmt"
	"sync"
)


func main(){
	var wg sync.WaitGroup

	incrementer := 0
	in := 50
	wg.Add(in)

	var mu sync.Mutex

// fmt.Println("CPUs:", runtime.NumCPU())
// fmt.Println(runtime.NumCgoCall())
// fmt.Println("Goroutines:", runtime.NumGoroutine())
	
	for i := 0 ; i < in; i++ {
		go func (){
				mu.Lock()
				v := incrementer
				v++
				incrementer = v
				fmt.Println(incrementer)
				mu.Unlock()
				wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementer)
}