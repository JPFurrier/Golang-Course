package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("This is a func expression")
	}

	f()

	g := func() {
		for i := 10; i >= 0; i-- {
			fmt.Println(i)
		}
	}
	fmt.Printf("%T\n", g)
	g()
}
