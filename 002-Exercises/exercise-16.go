package main

import "fmt"

func main() {

	x := make([]string, 50, 50)
	fmt.Println(len(x))

	x = []string{"Alabama", "Alasca", "Arkansas", "California", "Colorado", "Connecticut",
		"Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana",
		"Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan",
		"Minessota", "Mississipi", "Misscuri", "Montana", "Nebraska", "Nevada",
		"New Hempshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota",
		"Ohio", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina",
		"South Dakota", "Tehnessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wiscosin", "Wyoming"}

	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

}
