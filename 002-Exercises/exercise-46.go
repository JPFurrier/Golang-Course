package main

import "fmt"

func main() {

	c := make(chan int)

	//anonymous func
	go func() {
		c <- 42

	}()
	fmt.Println("This is a self-executing func", <-c)

	//buffered channel

	d := make(chan int, 2)

	d <- 50
	d <- 51

	fmt.Println("This is the buffered channel", <-d)
	fmt.Println("This is another value in bufferd channel", <-d)
}
