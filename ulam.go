package main

import (
	"flag"
	"fmt"
	"gerymate/ulam/sieve"
	"gerymate/ulam/spiral"
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
	u := spiral.NewSpiral(*n, s.IsPrime)
	u.FillSpiral()
	fmt.Println(u)
}
