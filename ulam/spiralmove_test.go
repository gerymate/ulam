package ulam

import (
	"testing"
)

func TestSpiralMove_NextMove(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		x int
		y int
		// Number of calls
		n     int
		wantx int
		wanty int
	}{
		{"From 2, 2, 1 times", 2, 2, 1, 3, 2},
		{"From 2, 2, 2 times", 2, 2, 2, 3, 1},
		{"From 2, 2, 3 times", 2, 2, 3, 2, 1},
		{"From 2, 2, 4 times", 2, 2, 4, 1, 1},
		{"From 2, 2, 5 times", 2, 2, 5, 1, 2},
		{"From 2, 2, 6 times", 2, 2, 6, 1, 3},
		{"From 2, 2, 7 times", 2, 2, 7, 2, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSpiralMove(tt.x, tt.y)
			for range tt.n - 1 {
				s.NextMove()
			}
			gotx, goty := s.NextMove()
			// TODO: update the condition below to compare got with tt.want.
			if gotx != tt.wantx || goty != tt.wanty {
				t.Errorf("NextMove() = %v, %v, want x: %v y: %v", gotx, goty, tt.wantx, tt.wanty)
			}
		})
	}
}
