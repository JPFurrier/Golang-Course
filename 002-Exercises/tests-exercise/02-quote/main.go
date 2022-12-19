package main

import (
	"fmt"
	"go-programming-master/code_samples/013-new-version/002-Exercises/tests-exercise/02-quote/quote"
	"go-programming-master/code_samples/013-new-version/002-Exercises/tests-exercise/02-quote/word"
)

func main() {
	fmt.Println("The string has exactly", word.Count(quote.SunAlso), "words!")

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
