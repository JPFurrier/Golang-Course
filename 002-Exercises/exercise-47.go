package main

import "fmt"

func main() {
	cs := make(chan int)

	//chan<- ------ send channel
	//<-chan ------ receive channel

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Println("---------")
	fmt.Printf("cs\t%T\n", cs)
}
