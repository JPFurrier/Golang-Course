package main

import "fmt"

func main() {

	// one way to make channel run

	//c := make(chan int)
	//go func() {c <- 42}()

	// another way to make channel run - buffer

	c := make(chan int, 1)
	c <- 42

	fmt.Println(<-c)
}
