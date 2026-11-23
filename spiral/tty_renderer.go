package spiral

import (
	"math"
	"strconv"
	"strings"
)

// Highlightmode
const (
	Mark = iota
	Invert
	Erase
	EraseOthers
)

type TTYRendererConfig struct {
	HighlightMode int
	OpenMark, CloseMark string
}

type TTYRenderer struct {
	selector func(int) bool
	dataSource *Spiral
	config TTYRendererConfig
}

func NewTTYRenderer(selector func(int)bool, dataSource *Spiral) *TTYRenderer {
	rc := TTYRendererConfig{
		HighlightMode: Mark,
		OpenMark: "(",
		CloseMark: ")",
	}
	return &TTYRenderer{selector, dataSource, rc}
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
			renderedNumber := formatNumber(number, numberWidth, selected, rr.config.OpenMark, rr.config.CloseMark, Mark)
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

func formatNumber(n, maxWidth int, mark bool, openM, closeM string, highlightMode int) string {
	open, close := " ", " "
	if mark {
		switch highlightMode {
		case Mark:
			open = openM
			close = closeM
		case Invert:
			open = "\x1b[7m"
			close = "\033[27m"
		case Erase:
			return strings.Repeat(" ", maxWidth)
		}
	} else if highlightMode == EraseOthers {
		return strings.Repeat(" ", maxWidth)
	}
	return open + alignNumber(n, maxWidth) + close		
}

func alignNumber(n, width int) string {
	numStr := strconv.Itoa(n)
	numLen := len(numStr)
	post := (width - numLen) / 2
	pre := post + (width - numLen) % 2
	return strings.Repeat(" ", pre) + numStr + strings.Repeat(" ", post)
}
