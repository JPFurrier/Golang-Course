package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {

	u1 := user{
		First: "James",
		Age:   23,
	}

	u2 := user{
		First: "Monneypenny",
		Age:   25,
	}

	u3 := user{
		First: "Dr. Who",
		Age:   99,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
