package spiral

import (
	"gerymate/ulam/sieve"
	"testing"
)

func TestSpiral_NewSpiral(t *testing.T) {
	cases := []struct {
		// in: numbers to make space for
		in   int
		want [][]int
	}{
		{1, [][]int{{0}}},
		{2, [][]int{{0, 0}, {0, 0}}},
		{3, [][]int{{0, 0}, {0, 0}}},
		{4, [][]int{{0, 0}, {0, 0}}},
		{5, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		{6, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
	}
	for _, c := range cases {
		selector := func(n int) bool {
			return false
		}
		got := NewSpiral(c.in, selector).numbers
		if !equalMatrices(got, c.want) {
			t.Errorf("NewSpiral(%d)==%d, want: %d", c.in, got, c.want)
		}
	}
}

func TestSpiral_FillSpiral(t *testing.T) {
	tests := []struct {
		name string
		// Fill the spiral up to (cn).
		cn int
		// Wanted output
		want [][]int
	}{
		{"1", 1, [][]int{{1}}},
		{"2", 2, [][]int{{0, 0}, {1, 2}}},
		{"3", 3, [][]int{{0, 3}, {1, 2}}},
		{"4", 4, [][]int{{4, 3}, {1, 2}}},
		{"5", 5, [][]int{{5, 4, 3}, {0, 1, 2}, {0, 0, 0}}},
		{"9", 9, [][]int{{5, 4, 3}, {6, 1, 2}, {7, 8, 9}}},
		{"10", 10, [][]int{{0, 0, 0, 0}, {5, 4, 3, 0}, {6, 1, 2, 0}, {7, 8, 9, 10}}},
		{"11", 11, [][]int{{0, 0, 0, 0}, {5, 4, 3, 0}, {6, 1, 2, 11}, {7, 8, 9, 10}}},
		{"13", 13, [][]int{{0, 0, 0, 13}, {5, 4, 3, 12}, {6, 1, 2, 11}, {7, 8, 9, 10}}},
		{"16", 16, [][]int{{16, 15, 14, 13}, {5, 4, 3, 12}, {6, 1, 2, 11}, {7, 8, 9, 10}}},
		{"17", 17, [][]int{{17, 16, 15, 14, 13}, {0, 5, 4, 3, 12}, {0, 6, 1, 2, 11}, {0, 7, 8, 9, 10}, {0, 0, 0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selector := func(n int) bool {
				return false
			}
			sp := NewSpiral(tt.cn, selector)
			sp.FillSpiral()
			got := sp.numbers
			if !equalMatrices(got, tt.want) {
				t.Errorf("FillSpiral on Spiral(%d)==%d, want: %d", tt.cn, got, tt.want)
			}
		})
	}
}

func Test_format_number_no_mark(t *testing.T) {
	tests := []struct {
		name string
		// Named input parameters for target function.
		n     int
		width int
		want  string
	}{
		{"1", 1, 1, " 1 "},
		{" 5", 5, 2, "  5 "},
		{"35", 35, 2, " 35 "},
		{" 8 ", 8, 3, "  8  "},
		{" 81", 81, 3, "  81 "},
		{"381", 381, 3, " 381 "},
		{"  4 ", 4, 4, "   4  "},
		{" 14 ", 14, 4, "  14  "},
		{" 147", 147, 4, "  147 "},
		{"5147", 5147, 4, " 5147 "},
		{"  2  ", 2, 5, "   2   "},
		{"  23 ", 23, 5, "   23  "},
		{" 123 ", 123, 5, "  123  "},
		{" 1234", 1234, 5, "  1234 "},
		{"91234", 91234, 5, " 91234 "},
		{"   7  ", 7, 6, "    7   "},
		{"  17  ", 17, 6, "   17   "},
		{"  817 ", 817, 6, "   817  "},
		{" 2817 ", 2817, 6, "  2817  "},
		{" 28179", 28179, 6, "  28179 "},
		{"328179", 328179, 6, " 328179 "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatNumber(tt.n, tt.width, false)
			if got != tt.want {
				t.Errorf("formatNumber() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_format_number_mark(t *testing.T) {
	tests := []struct {
		name string
		// Named input parameters for target function.
		n     int
		width int
		want  string
	}{
		{"1", 1, 1, "(1)"},
		{" 5", 5, 2, " (5)"},
		{"35", 35, 2, "(35)"},
		{" 8 ", 8, 3, " (8) "},
		{" 81", 81, 3, " (81)"},
		{"381", 381, 3, "(381)"},
		{"  4 ", 4, 4, "  (4) "},
		{" 14 ", 14, 4, " (14) "},
		{" 147", 147, 4, " (147)"},
		{"5147", 5147, 4, "(5147)"},
		{"  2  ", 2, 5, "  (2)  "},
		{"  23 ", 23, 5, "  (23) "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatNumber(tt.n, tt.width, true)
			if got != tt.want {
				t.Errorf("formatNumber() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestSpiral_String_Integers(t *testing.T) {
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
			sp := NewSpiral(tt.n, selector)
			sp.FillSpiral()
			got := sp.String()
			if got != tt.want {
				t.Errorf("String() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}

func TestSpiral_String_Primality(t *testing.T) {
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
			sp := NewSpiral(tt.n, selector)
			sp.FillSpiral()
			got := sp.String()
			if got != tt.want {
				t.Errorf("String() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}
