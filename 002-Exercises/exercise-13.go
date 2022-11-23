package main

import "fmt"

func main() {

	// array

	x := [5]int{1, 2, 3, 4, 5}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

	// slice
	
	y := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i, v := range y {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", y)
}

