package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin goroutine", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello From the Other Side")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from this side")
		wg.Done()
	}()

	fmt.Println("MID cpu", runtime.NumCPU())
	fmt.Println("MID goroutine", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end goroutine", runtime.NumGoroutine())
}
