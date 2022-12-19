package word

import (
	"fmt"
	"go-programming-master/code_samples/013-new-version/002-Exercises/tests-exercise/02-quote/quote"
	"testing"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("Got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("Got", v, "want", 3)
			}
		}
	}
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("Expected", 3, "Got", n)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	//Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}

}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
