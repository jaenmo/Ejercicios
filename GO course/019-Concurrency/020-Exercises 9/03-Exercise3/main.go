package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main(){
	var wg sync.WaitGroup

	incrementer := 0
	in := 40
	wg.Add(in)

// fmt.Println("CPUs:", runtime.NumCPU())
// fmt.Println(runtime.NumCgoCall())
// fmt.Println("Goroutines:", runtime.NumGoroutine())
	
	for i := 0 ; i < in; i++ {
		go func (){
			
				v := incrementer
				runtime.Gosched()
				v++
				incrementer = v
				fmt.Println(incrementer)
				wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementer)
}