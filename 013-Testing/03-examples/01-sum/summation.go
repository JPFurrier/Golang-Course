// Package summation take the values and returns the sum
package summation

// Sum adds an unlimited number of values of type int
func Sum(xi ...int) int {
	s := 0

	for _, v := range xi {
		s += v
	}
	return s
}
