package main

import "fmt"

// slice allows you to group together VALUES of the same type

func main() {

	// x := type{values} -> COMPOSITE LITERAL
	//	x := []int{4, 5, 6, 7, 20, 88, 77}
	//	fmt.Println(x)

	// for range over slice

	//x := []int{1, 6, 5, 4, 33, 11, 67}
	//for i, v := range x {
	//	fmt.Println(i, v)
	//}

	// slicing a slice
	//x := []int{5, 6, 7, 8, 9, 10}
	//fmt.Println(x)
	//fmt.Println(x[1:])
	//fmt.Println(x[2:6])
	//
	//for i := 0; i <= 5; i++ {
	//	fmt.Println(i, x[i])
	//}

	// append a slice
	//x := []int{4, 5, 6, 7, 8, 9}
	//fmt.Println(x)
	//x = append(x, 22, 33, 44, 55, 66, 77)
	//fmt.Println(x)
	//
	//y := []int{123, 456, 789}
	//x = append(x, y...)
	//fmt.Println(y)
	//fmt.Println(x)
	//
	//// delete slice
	//x = append(x[8:10], x[13:]...)
	//fmt.Println(x)

	//multi-dimensional slice
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
