// Write the numbers in a spiral pattern
package ulam

import (
	"math"
	"strconv"
	"strings"
)

type Ulam struct {
	// Will hold the numbers up to (n) arranged in a spiral.
	selector func(int)bool
	n       int
	numbers [][]int
}

func NewUlam(n int, selector func(int)bool) *Ulam {
	// Constructor for Ulam object. Makes the required space.
	dim := int(math.Ceil(math.Sqrt(float64(n))))

	myUlam := Ulam{n: n, selector: selector}
	myUlam.numbers = make([][]int, dim)
	for i := range dim {
		myUlam.numbers[i] = make([]int, dim)
	}
	return &myUlam
}

func (u *Ulam) FillUlam() {
	// Fill the Ulam object with the natural numbers in a spiral.
	dim := len(u.numbers)
	x := (dim - 1) / 2
	y := x + (dim-1)%2
	mover := NewSpiralMove(x, y)
	for i := range u.n {
		u.numbers[y][x] = i + 1
		x, y = mover.NextMove()
	}
}

func (u *Ulam) String() string {
	// Render the Ulam object.
	number_width := 1 + int(math.Log10(float64(u.n)))
	lines := make([]string, 0, len(u.numbers))
	for _, row_of_numbers := range(u.numbers) {
		row_of_strings := make([]string, 0, len(row_of_numbers))
		// Render each number. Use spaces if 0. Skip line if all 0.
		skip := true
		for _, number := range(row_of_numbers) {
			// mark if it's important
			selected := u.selector(number)
			// use our formatter
			rendered_number := format_number(number, number_width, selected)
			// handle special cases, ie. when there's nothing to render
			if number == 0 {
				rendered_number = strings.Repeat(" ", number_width + 2)
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
	pre := post + (width - num_len) % 2
	return strings.Repeat(" ", pre) + open + num_str + close + strings.Repeat(" ", post)
}

