package spiral

import (
	"math"
	"strconv"
	"strings"
)

type HlMode int

// Highlightmode
const (
	Mark HlMode = iota
	Invert
	Erase
	EraseOthers
)

type TTYRendererConfig struct {
	HighlightMode       HlMode
	OpenMark, CloseMark string
}

type TTYRenderer struct {
	selector   func(int) bool
	dataSource *Spiral
	Config     TTYRendererConfig
}

func NewTTYRenderer(selector func(int) bool, dataSource *Spiral, config TTYRendererConfig) *TTYRenderer {
	return &TTYRenderer{selector, dataSource, config}
}

func DefaultTTYRenderer(selector func(int) bool, dataSource *Spiral) *TTYRenderer {
	rc := TTYRendererConfig{
		HighlightMode: Mark,
		OpenMark:      "(",
		CloseMark:     ")",
	}
	return NewTTYRenderer(selector, dataSource, rc)
}

func (rr *TTYRenderer) Render() string {
	// Render the Spiral object.
	numberWidth := 1 + int(math.Log10(float64(rr.dataSource.n)))
	lines := make([]string, 0, len(rr.dataSource.numbers))
	for _, rowOfNumbers := range rr.dataSource.numbers {
		rowOfStrings := make([]string, 0, len(rowOfNumbers))
		// Render each number. Mark selected ones. Skip line if all 0.
		skip := true
		for _, number := range rowOfNumbers {
			selected := rr.selector(number)
			renderedNumber := formatNumber(number, numberWidth, selected, rr.Config.OpenMark, rr.Config.CloseMark, rr.Config.HighlightMode)
			if number != 0 {
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

func formatNumber(n, maxWidth int, mark bool, openM, closeM string, highlightMode HlMode) string {
	open, close := " ", " "
	if highlightMode == Invert {
		open, close = "", ""
	}
	if mark {
		switch highlightMode {
		case Mark:
			open = openM
			close = closeM
		case Invert:
			open = "\x1b[7m"
			close = "\033[27m"
		case Erase:
			return strings.Repeat(" ", maxWidth + 2)
		}
	} else if highlightMode == EraseOthers {
		return strings.Repeat(" ", maxWidth + 2)
	}
	if n == 0 {
		if highlightMode == Invert {
			return strings.Repeat(" ", maxWidth)
		}
		return strings.Repeat(" ", maxWidth + 2)
	}
	return open + alignNumber(n, maxWidth) + close
}

func alignNumber(n, width int) string {
	numStr := strconv.Itoa(n)
	numLen := len(numStr)
	post := (width - numLen) / 2
	pre := post + (width - numLen ) % 2
	return strings.Repeat(" ", pre) + numStr + strings.Repeat(" ", post)
}
