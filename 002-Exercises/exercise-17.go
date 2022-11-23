package main

import "fmt"

func main() {

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Hellooooo, James"}
	fmt.Println(mp)

	jp := [][]string{jb, mp}
	fmt.Print(jp)

	for i, v := range jp {
		fmt.Println("Record:", i)
		for j, val := range v {
			fmt.Printf("\t index position %v \t value: \t %v", j, val)
		}
	}

}
