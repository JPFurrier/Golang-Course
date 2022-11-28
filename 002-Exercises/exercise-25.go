package main

import "fmt"

func main() {
	defer fmt.Println("This is after, 'cause its was defer")
	fmt.Println("This will run first")

	foo()
}

func foo() {
	defer func() {
		fmt.Println("This anonymous func was DEFER")
	}()
	fmt.Println("This is foo func")
}
