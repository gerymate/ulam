package main

import (
	"fmt"
	"gerymate/ulam/ulam"
)

func main() {
	u := ulam.NewUlam(141)
	u.FillUlam()
	fmt.Println(u)
}
