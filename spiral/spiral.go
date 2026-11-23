// Write the numbers in a spiral pattern
package spiral

import (
	"math"
)

type Spiral struct {
	// Will hold the numbers up to (n) arranged in a spiral.
	n        int
	numbers  [][]int
}

func NewSpiral(n int) *Spiral {
	// Constructor for Spiral object. Makes the required space.
	dim := int(math.Ceil(math.Sqrt(float64(n))))

	mySpiral := Spiral{n: n}
	mySpiral.numbers = make([][]int, dim)
	for i := range dim {
		mySpiral.numbers[i] = make([]int, dim)
	}
	return &mySpiral
}

func (sp *Spiral) FillSpiral() {
	// Fill the Spiral object with the natural numbers.
	dim := len(sp.numbers)
	x := (dim - 1) / 2
	y := x + (dim - 1) % 2
	mover := NewSpiralMove(x, y)
	for i := range sp.n {
		sp.numbers[y][x] = i + 1
		x, y = mover.NextMove()
	}
}

