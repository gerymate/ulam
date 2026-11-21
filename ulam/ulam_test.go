package ulam

import (
	"testing"
)

func TestUlam_NewUlam(t *testing.T) {
	cases := []struct {
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
		got := NewUlam(c.in).numbers
		if !equalMatrices(got, c.want) {
			t.Errorf("NewUlam(%d)==%d, want: %d", c.in, got, c.want)
		}
	}
}

func TestUlam_FillUlam(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
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
			u := NewUlam(tt.cn)
			u.FillUlam()
			got := u.numbers
			if !equalMatrices(got, tt.want) {
				t.Errorf("FillUlam on Ulam(%d)==%d, want: %d", tt.cn, got, tt.want)
			}
		})
	}
}

func TestUlam_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		n    int
		want string
	}{
		{"1", 1, "1"},
		{"2", 2, "1 2"},
		{"3", 3, "  3\n1 2"},
		{"4", 4, "4 3\n1 2"},
		{"5", 5, "5 4 3\n  1 2"},
		{"6", 6, "5 4 3\n6 1 2"},
		{"7", 7, "5 4 3\n6 1 2\n7    "},
		{"8", 8, "5 4 3\n6 1 2\n7 8  "},
		{"9", 9, "5 4 3\n6 1 2\n7 8 9"},
		{"15", 15, "   15 14 13\n 5  4  3 12\n 6  1  2 11\n 7  8  9 10"},
		{"17", 17, "17 16 15 14 13\n    5  4  3 12\n    6  1  2 11\n    7  8  9 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUlam(tt.n)
			u.FillUlam()
			got := u.String()
			if got != tt.want {
				t.Errorf("String() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}
