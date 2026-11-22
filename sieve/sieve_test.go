package sieve

import (
	"slices"
	"strconv"
	"testing"
)

func TestSieve_IsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
		{17, true},
		{18, false},
		{19, true},
		{20, false},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			s := NewSieve(tt.n)
			got := s.IsPrime(tt.n)
			if got != tt.want {
				t.Errorf("IsPrime(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func Test_removeNumber(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr    []bool
		number int
		want   []bool
	}{
		{"3s", []bool{true, true, true, true, true, true, true, true}, 3, []bool{true, true, true, false, true, true, false, true}},
		{"5s", []bool{true, true, true, true, true, true, true, true}, 5, []bool{true, true, true, true, true, true, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNumber(tt.arr, tt.number)
			if !slices.Equal(tt.want, got) {
				t.Errorf("removeNumber(%v) = %v, want %v", tt.number, got, tt.want)
			}
		})
	}
}
