package main

import (
	"fmt"
	"go-programming-master/code_samples/013-new-version/012-Documentation/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(3),
	}
	fmt.Println(fido)
}
