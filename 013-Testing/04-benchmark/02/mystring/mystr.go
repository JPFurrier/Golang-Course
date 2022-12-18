package mystr

import "strings"

// Concatenate the text of a string
// Low efficient and hard processing
func Cat(xs []string) string {
	s := ""

	for _, v := range xs {
		s += v
		s += " "
	}
	return s
}

// concatenate the text of a string
// More efficient
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
