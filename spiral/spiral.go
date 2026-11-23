// Write the numbers in a spiral pattern
package spiral

import (
	"math"
	"strconv"
	"strings"
)

type Spiral struct {
	// Will hold the numbers up to (n) arranged in a spiral.
	selector func(int) bool
	n        int
	numbers  [][]int
}

func NewSpiral(n int, selector func(int) bool) *Spiral {
	// Constructor for Spiral object. Makes the required space.
	dim := int(math.Ceil(math.Sqrt(float64(n))))

	mySpiral := Spiral{n: n, selector: selector}
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
	y := x + (dim-1)%2
	mover := NewSpiralMove(x, y)
	for i := range sp.n {
		sp.numbers[y][x] = i + 1
		x, y = mover.NextMove()
	}
}

func (sp *Spiral) String() string {
	// Render the Spiral object.
	numberWidth := 1 + int(math.Log10(float64(sp.n)))
	lines := make([]string, 0, len(sp.numbers))
	for _, rowOfNumbers := range sp.numbers {
		rowOfStrings := make([]string, 0, len(rowOfNumbers))
		// Render each number. Use spaces if 0. Skip line if all 0.
		skip := true
		for _, number := range rowOfNumbers {
			// mark if it's important
			selected := sp.selector(number)
			// use our formatter
			renderedNumber := formatNumber(number, numberWidth, selected)
			// handle special cases, ie. when there's nothing to render
			if number == 0 {
				renderedNumber = strings.Repeat(" ", numberWidth + 2)
			} else {
				skip = false
			}
			rowOfStrings = append(rowOfStrings, renderedNumber)
		}
		if skip {
			continue
		}
		lines = append(lines, strings.Join(rowOfStrings, " "))
	}
	return strings.Join(lines, "\n")
}

func formatNumber(n, width int, mark bool) string {
	numStr := strconv.Itoa(n)
	numLen := len(numStr)
	open, close := " ", " "
	if mark {
		open = "("
		close = ")"
	}
	post := (width - numLen) / 2
	pre := post + (width - numLen) % 2
	return strings.Repeat(" ", pre) + open + numStr + close + strings.Repeat(" ", post)
}
