package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6}
	n := foo(s...)
	fmt.Println(n)

	f := []int{1, 2, 3, 4, 5, 6}
	g := bar(f)
	fmt.Println(g)

}

// variadic parameter
func foo(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}
