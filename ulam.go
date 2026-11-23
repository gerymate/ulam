package main

import (
	"flag"
	"fmt"
	"gerymate/ulam/sieve"
	"gerymate/ulam/spiral"
)

func main() {
	n := flag.Int("n", 150, "generate Ulam's spiral up to n")
	h := flag.Bool("h", false, "get help")
	m1 := flag.String("m1", "[", "opening mark")
	m2 := flag.String("m2", "]", "closing mark")

	hlm := spiral.Invert
	flag.Func("hlm", "highlight mode: mark|invert|erase|eraseothers (default invert)", func(s string) error {
		switch s {
		case "mark":
			hlm = spiral.Mark
		case "invert":
			hlm = spiral.Invert
		case "erase":
			hlm = spiral.Erase
		case "eraseothers":
			hlm = spiral.EraseOthers
		default:
			return fmt.Errorf("unknown highlight mode: %q", s)
		}
		return nil
	})

	flag.Parse()
	if *h {
		flag.PrintDefaults()
		return
	}
	sv := sieve.NewSieve(*n)
	sp := spiral.NewSpiral(*n)
	sp.FillSpiral()
	config := spiral.TTYRendererConfig{
		HighlightMode: hlm,
		OpenMark:      *m1,
		CloseMark:     *m2,
	}
	rr := spiral.NewTTYRenderer(sv.IsPrime, sp, config)

	fmt.Println(rr.Render())
}
