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
	number_width := 1 + int(math.Log10(float64(sp.n)))
	lines := make([]string, 0, len(sp.numbers))
	for _, row_of_numbers := range sp.numbers {
		row_of_strings := make([]string, 0, len(row_of_numbers))
		// Render each number. Use spaces if 0. Skip line if all 0.
		skip := true
		for _, number := range row_of_numbers {
			// mark if it's important
			selected := sp.selector(number)
			// use our formatter
			rendered_number := format_number(number, number_width, selected)
			// handle special cases, ie. when there's nothing to render
			if number == 0 {
				rendered_number = strings.Repeat(" ", number_width+2)
			} else {
				skip = false
			}
			row_of_strings = append(row_of_strings, rendered_number)
		}
		if skip {
			continue
		}
		lines = append(lines, strings.Join(row_of_strings, " "))
	}
	return strings.Join(lines, "\n")
}

func format_number(n, width int, mark bool) string {
	num_str := strconv.Itoa(n)
	num_len := len(num_str)
	open, close := " ", " "
	if mark {
		open = "("
		close = ")"
	}
	post := (width - num_len) / 2
	pre := post + (width-num_len)%2
	return strings.Repeat(" ", pre) + open + num_str + close + strings.Repeat(" ", post)
}
