package main

// fixing the race condition with func mutex

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementer)
}
