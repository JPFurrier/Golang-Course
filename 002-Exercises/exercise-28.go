package main

import "fmt"

func main() {

	//anonymous func
	func() {
		fmt.Println("This is the first anonymous func")
	}()

	func(x int) {
		fmt.Println("The result of anonymous func is:", x)
	}(42)

	func() {
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("Done!")

}
