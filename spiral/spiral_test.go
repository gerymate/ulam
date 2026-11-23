package spiral

import (
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
		got := NewSpiral(c.in).numbers
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
			sp := NewSpiral(tt.cn)
			sp.FillSpiral()
			got := sp.numbers
			if !equalMatrices(got, tt.want) {
				t.Errorf("FillSpiral on Spiral(%d)==%d, want: %d", tt.cn, got, tt.want)
			}
		})
	}
}
