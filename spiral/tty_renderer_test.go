package spiral

import (
	"gerymate/ulam/sieve"
	"testing"
)

func TestTTYRenderer_Render_Integers(t *testing.T) {
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
			rr := DefaultTTYRenderer(selector, sp)
			got := rr.Render()
			if got != tt.want {
				t.Errorf("Render() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}

func TestTTYRenderer_Render_Primality(t *testing.T) {
	tests := []struct {
		name string
		// Named input parameters for receiver constructor.
		n    int
		want string
	}{
		{"3", 3, "    (3)\n 1  (2)"},
		{"4", 4, " 4  (3)\n 1  (2)"},
		{"8", 8, "(5)  4  (3)\n 6   1  (2)\n(7)  8     "},
		{"15", 15, "      15   14  (13)\n( 5)   4  ( 3)  12 \n  6    1  ( 2) (11)\n( 7)   8    9   10 "},
		{"17", 17, "(17)  16   15   14  (13)\n     ( 5)   4  ( 3)  12 \n       6    1  ( 2) (11)\n     ( 7)   8    9   10 "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sieve.NewSieve(tt.n)
			selector := s.IsPrime
			sp := NewSpiral(tt.n)
			sp.FillSpiral()
			rr := DefaultTTYRenderer(selector, sp)
			got := rr.Render()
			if got != tt.want {
				t.Errorf("String() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}

func Test_formatNumber_no_mark(t *testing.T) {
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
			got := formatNumber(tt.n, tt.width, false, "(", ")", Mark)
			if got != tt.want {
				t.Errorf("formatNumber() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_formatNumber_mark(t *testing.T) {
	tests := []struct {
		name string
		// Named input parameters for target function.
		n     int
		width int
		want  string
	}{
		{"1", 1, 1, "(1)"},
		{" 5", 5, 2, "( 5)"},
		{"35", 35, 2, "(35)"},
		{" 8 ", 8, 3, "( 8 )"},
		{" 81", 81, 3, "( 81)"},
		{"381", 381, 3, "(381)"},
		{"  4 ", 4, 4, "(  4 )"},
		{" 14 ", 14, 4, "( 14 )"},
		{" 147", 147, 4, "( 147)"},
		{"5147", 5147, 4, "(5147)"},
		{"  2  ", 2, 5, "(  2  )"},
		{"  23 ", 23, 5, "(  23 )"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatNumber(tt.n, tt.width, true, "(", ")", Mark)
			if got != tt.want {
				t.Errorf("formatNumber() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_formatNumber(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n             int
		width         int
		mark          bool
		openM         string
		closeM        string
		highlightMode HlMode
		want          string
	}{
		{"*8*", 8, 1, true, "*", "*", Mark, "*8*"},
		{"[7]", 7, 1, true, "[", "]", Mark, "[7]"},
		{"Invert 6", 6, 3, true, "", "", Invert, "\x1b[7m 6 \033[27m"},
		{"Erase 5", 5, 3, true, "", "", Erase, "     "},
		{"EraseOthers 4", 4, 2, false, "", "", EraseOthers, "    "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatNumber(tt.n, tt.width, tt.mark, tt.openM, tt.closeM, tt.highlightMode)
			if got != tt.want {
				t.Errorf("formatNumber() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestTTYRenderer_Render(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		n   int
		config     TTYRendererConfig
		want       string
	}{
		{"Mark", 3, TTYRendererConfig{Mark, "{", "}"}, "    {3}\n 1  {2}"},
		{"Erase", 3, TTYRendererConfig{Erase, "/", "|"}, "       \n 1     "},
		{"EraseOthers", 3, TTYRendererConfig{EraseOthers, "", ""}, "     3 \n     2 "},
		{"Invert", 3, TTYRendererConfig{Invert, " ", " "}, "  \x1b[7m3\033[27m\n1 \x1b[7m2\033[27m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sieve.NewSieve(tt.n)
			selector := s.IsPrime
			sp := NewSpiral(tt.n)
			sp.FillSpiral()
			rr := NewTTYRenderer(selector, sp, tt.config)
			got := rr.Render()
			if got != tt.want {
				t.Errorf("Render() = \n'%v'\n, want \n'%v'", got, tt.want)
			}
		})
	}
}
