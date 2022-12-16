package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtErr struct {
	lat  string
	long string
	err  error
}

func (se sqrtErr) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := errors.New("more coffee needed")
		return 0, sqrtErr{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
