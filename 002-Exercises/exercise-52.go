package main

import "fmt"

func main() {

	c := make(chan int)

	//lauches 10 goroutines and each goroutine add 10 numbers to a channel

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}
