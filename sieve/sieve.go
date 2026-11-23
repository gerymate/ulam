package sieve

import "math"

type Sieve struct {
	// This Sieve will store primality for odd numbers beginning with 3
	primes []bool
}

func NewSieve(n int) *Sieve {
	// Prepare a Sieve
	s := Sieve{}
	if n < 3 {
		return &s
	}
	s.primes = make([]bool, n / 2)
	s.fillSieve(n)
	return &s
}

func (s *Sieve)fillSieve(n int) {
	primes := s.primes
	for i := range(primes) {
		primes[i] = true
	}
	end := int(math.Sqrt(float64(n)))
	end_index := (end - 3 / 2)
	for index := 0; index <= end_index; index++ {
		if !primes[index] {
			continue
		}
		primes = removeNumber(primes, (2 * index) + 3)
	}
	s.primes = primes
}

func removeNumber(arr []bool, number int) []bool {
	start := (number - 3) / 2
	for i := start + number; i < len(arr); i += number {
		arr[i] = false
	}
	return arr
}

func (s *Sieve)IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	} else if n % 2 == 0 {
		return false
	}

	return s.primes[(n - 3) / 2]
}