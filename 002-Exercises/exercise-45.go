package main

import (
	"fmt"
	"runtime"
)

// prints out your OS and ARCH

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
