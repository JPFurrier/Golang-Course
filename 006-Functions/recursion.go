package main

import "fmt"

func main() {

	m := factorial(4)
	fmt.Println(m)

	n := loopFact(4)
	fmt.Println(n)
}

// recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// using loop

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
