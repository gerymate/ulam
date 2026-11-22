package main

import (
	"flag"
	"fmt"
	"gerymate/ulam/sieve"
	"gerymate/ulam/ulam"
)

func main() {
	n := flag.Int("n", 200, "generate Ulam's spiral up to n")
	h := flag.Bool("h", false, "get help")
	flag.Parse()
	if *h {
		flag.PrintDefaults()
		return
	}
	s := sieve.NewSieve(*n)
	u := ulam.NewUlam(*n, s.IsPrime)
	u.FillUlam()
	fmt.Println(u)
}
