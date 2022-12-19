package dog

import (
	"fmt"
	"testing"
)

// Test the Years function
func TestYears(t *testing.T) {
	x := Years(3)
	if x != 21 {
		t.Error("Expected", 21, "Got", x)
	}
}

// Test the YearsTwo function
func TestYearsTwo(t *testing.T) {
	x := YearsTwo(3)
	if x != 21 {
		t.Error("Expected", 21, "Got", x)
	}
}

// Example
func ExampleYears() {
	fmt.Println(Years(2))
	//Output:
	//14
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(2))
	//Output:
	//14
}

// Benchmark Years
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

// Benchmark YearsTwo
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
