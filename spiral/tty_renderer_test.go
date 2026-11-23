package spiral

import (
	"gerymate/ulam/sieve"
	"testing"
)

func TestSpiral_Render_Integers(t *testing.T) {
	tests := []struct {
		name string
		// Generate string output for a spiral of natural numbers up to (n).
		n    int
		want string
	}{
		{"1", 1, " 1 "},
		{"2", 2, " 1   2 "},
		{"3", 3, "     3 \n 1   2 "},
		{"4", 4, " 4   3 \n 1   2 "},
		{"5", 5, " 5   4   3 \n     1   2 "},
		{"6", 6, " 5   4   3 \n 6   1   2 "},
		{"7", 7, " 5   4   3 \n 6   1   2 \n 7         "},
		{"8", 8, " 5   4   3 \n 6   1   2 \n 7   8     "},
		{"9", 9, " 5   4   3 \n 6   1   2 \n 7   8   9 "},
		{"15", 15, "      15   14   13 \n  5    4    3   12 \n  6    1    2   11 \n  7    8    9   10 "},
		{"17", 17, " 17   16   15   14   13 \n       5    4    3   12 \n       6    1    2   11 \n       7    8    9   10 "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selector := func(n int) bool {
				return false
			}
			sp := NewSpiral(tt.n)
			sp.FillSpiral()
			rr := NewTTYRenderer(selector, sp)
			got := rr.Render()
			if got != tt.want {
				t.Errorf("Render() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}

func TestSpiral_Render_Primality(t *testing.T) {
	tests := []struct {
		name string
		// Named input parameters for receiver constructor.
		n    int
		want string
	}{
		{"1", 1, " 1 "},
		{"2", 2, " 1  (2)"},
		{"3", 3, "    (3)\n 1  (2)"},
		{"4", 4, " 4  (3)\n 1  (2)"},
		{"5", 5, "(5)  4  (3)\n     1  (2)"},
		{"6", 6, "(5)  4  (3)\n 6   1  (2)"},
		{"7", 7, "(5)  4  (3)\n 6   1  (2)\n(7)        "},
		{"8", 8, "(5)  4  (3)\n 6   1  (2)\n(7)  8     "},
		{"9", 9, "(5)  4  (3)\n 6   1  (2)\n(7)  8   9 "},
		{"15", 15, "      15   14  (13)\n (5)   4   (3)  12 \n  6    1   (2) (11)\n (7)   8    9   10 "},
		{"17", 17, "(17)  16   15   14  (13)\n      (5)   4   (3)  12 \n       6    1   (2) (11)\n      (7)   8    9   10 "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sieve.NewSieve(tt.n)
			selector := s.IsPrime
			sp := NewSpiral(tt.n)
			sp.FillSpiral()
			rr := NewTTYRenderer(selector, sp)
			got := rr.Render()
			if got != tt.want {
				t.Errorf("String() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}
