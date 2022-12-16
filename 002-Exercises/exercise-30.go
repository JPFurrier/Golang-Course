package main

import "fmt"

func main() {

	//assign a func on a variable and print that out
	f := foo()
	fmt.Println(f())

}

// a func that return a func that return an int
func foo() func() int {
	return func() int {
		return 42
	}
}
