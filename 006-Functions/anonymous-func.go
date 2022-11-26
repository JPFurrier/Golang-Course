package main

import "fmt"

func main() {

	foo()

	//anonymous func
	func() {
		fmt.Println("Anonymous func ran")
	}()

	// passando argumentos para func anonima
	func(x int) {
		fmt.Println("The meaning of life is:", x)
	}(42)

}

func foo() {
	fmt.Println("foo ran")
}
