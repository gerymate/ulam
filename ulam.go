package main

import (
	"fmt"
	"gerymate/ulam/sieve"
	"gerymate/ulam/ulam"
)

func main() {
	n := 200
	s := sieve.NewSieve(n)
	u := ulam.NewUlam(n, s.IsPrime)
	u.FillUlam()
	fmt.Println(u)
}
