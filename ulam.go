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
	sv := sieve.NewSieve(*n)
	sp := spiral.NewSpiral(*n)
	sp.FillSpiral()
	rr := spiral.NewTTYRenderer(sv.IsPrime, sp)
	fmt.Println(rr.Render())
}
