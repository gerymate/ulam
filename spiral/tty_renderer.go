package spiral

import (
	"math"
	"strconv"
	"strings"
)

type TTYRenderer struct {
	selector func(int) bool
	dataSource *Spiral
}

func NewTTYRenderer(selector func(int)bool, dataSource *Spiral) *TTYRenderer {
	return &TTYRenderer{selector, dataSource}
}

func (rr *TTYRenderer) Render() string {
	// Render the Spiral object.
	numberWidth := 1 + int(math.Log10(float64(rr.dataSource.n)))
	lines := make([]string, 0, len(rr.dataSource.numbers))
	for _, rowOfNumbers := range rr.dataSource.numbers {
		rowOfStrings := make([]string, 0, len(rowOfNumbers))
		// Render each number. Use spaces if 0. Skip line if all 0.
		skip := true
		for _, number := range rowOfNumbers {
			// mark if it's important
			selected := rr.selector(number)
			// use our formatter
			renderedNumber := formatNumber(number, numberWidth, selected)
			// handle special cases, ie. when there's nothing to render
			if number == 0 {
				renderedNumber = strings.Repeat(" ", numberWidth+2)
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
	pre := post + (width-numLen)%2
	return strings.Repeat(" ", pre) + open + numStr + close + strings.Repeat(" ", post)
}
