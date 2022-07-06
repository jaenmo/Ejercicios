package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


func main(){

	var counter int64
	var wg sync.WaitGroup

	in := 50
	wg.Add(in)

// fmt.Println("CPUs:", runtime.NumCPU())
// fmt.Println(runtime.NumCgoCall())
// fmt.Println("Goroutines:", runtime.NumGoroutine())
	
	for i := 0 ; i < in; i++ {
		go func (){
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value", counter)
}