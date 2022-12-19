package main

import (
	"fmt"
	dog "go-programming-master/code_samples/013-new-version/002-Exercises/tests-exercise/01-dog/01"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
