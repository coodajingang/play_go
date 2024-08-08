package main

import "testing"

func main() {

}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// go test -v
func TestIntMinBasic(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{1, 0, 0},
		{0, 0, 0},
		{1, 2, 1},
		{-1, 2, -1},
		{-1, 2, -1},
	}

	for _, test := range tests {
		got := IntMin(test.a, test.b)
		if got != test.want {
			t.Errorf("IntMin(%d, %d) =%d, but want %d", test.a, test.b, got, test.want)
		}
	}
}

// go test -bench=.
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
