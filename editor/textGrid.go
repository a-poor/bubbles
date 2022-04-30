package editor

import (
	"fmt"
	"strings"
)

type textGrid struct {
	data [][]rune
}

// newTextGrid creates a new empty textGrid.
func newTextGrid() *textGrid {
	return &textGrid{
		data: make([][]rune, 0),
	}
}

// textGridFromString creates a new textGrid
// seeded with data in string s.
func textGridFromString(s string) *textGrid {
	lines := strings.Split(s, "\n")
	d := make([][]rune, len(lines), 0)
	for i, l := range lines {
		d[i] = []rune(l)
	}
	return &textGrid{
		data: d,
	}
}

// String converts a textGrid to its string
// representation.
func (tg *textGrid) String() string {
	var lines []string
	for _, l := range tg.data {
		lines = append(lines, string(l))
	}
	return strings.Join(lines, "\n")
}

// length returns the number of lines in the textGrid,
func (tg *textGrid) length() int {
	return len(tg.data)
}

// widthAt returns the number of runes in the i-th
// line of the textGrid.
//
// Note: If i is out of bounds, widthAt will panic.
func (tg *textGrid) widthAt(i int) int {
	if i < 0 || i >= len(tg.data) {
		err := fmt.Errorf("index %d out of bounds for grid of length %d", i, len(tg.data))
		panic(err)
	}
	return len(tg.data[i])
}

// addLineAt adds a new empty line at the i-th line
// index.
//
// Note: If i is out of bounds, widthAt will panic.
func (tg *textGrid) addLineAt(i int) {
	// Bounds check the input
	n := tg.length()
	if i < 0 || i > n {
		err := fmt.Errorf("index %d out of bounds for grid of length %d", i, n)
		panic(err)
	}

	// Create a new slice to store the output
	newData := make([][]rune, n+1)

	// Add the data before i
	for idx := 0; idx < i; idx++ {
		newData[idx] = tg.data[idx]
	}

	// Add an empty line at i
	newData[i] = make([]rune, 0)

	// Add the data after i
	for idx := i; idx < n; idx++ {
		newData[idx+1] = tg.data[idx]
	}

	// Set the result
	tg.data = newData
}

// appendLine adds a new empty line at the end of
// the textGrid.
func (tg *textGrid) appendLine() {
	n := tg.length()
	tg.addLineAt(n)
}

// appendLine adds a new empty line at the beginning
// of the textGrid.
func (tg *textGrid) prependLine() {
	tg.addLineAt(0)
}

// getLine returns the value of the i-th line in
// the textGrid as a string.
//
// Note: If i is out of bounds, widthAt will panic.
func (tg *textGrid) getLine(i int) string {
	// Bounds check the input
	n := tg.length()
	if i < 0 || i >= n {
		err := fmt.Errorf("index %d out of bounds for grid of length %d", i, n)
		panic(err)
	}
	return string(tg.data[i])
}

// getLines returns a slice of all lines in the
// textGrid.
func (tg *textGrid) getLines() []string {
	var lines []string
	for _, l := range tg.data {
		lines = append(lines, string(l))
	}
	return lines
}

// setLine sets the value of the i-th line of the
// textGrid to the value of line.
func (tg *textGrid) setLine(i int, line string) {
	// Bounds check the input
	n := tg.length()
	if i < 0 || i >= n {
		err := fmt.Errorf("index %d out of bounds for grid of length %d", i, n)
		panic(err)
	}
	tg.data[i] = []rune(line)
}

// splitLineAt creates a line break at the j-th
// character on the i-th line.
//
// Any characters before the j-th character will
// remain on the i-th line. Any subsequent characters
// will be moved to a new line at i+1.
//
// Note: If i or j is out of bounds, splitLineAt will
// panic. The value of i can be in the range
// [0, number_of_lines) while the value of j can be
// in the range [0, width_at_line_i].
func (tg *textGrid) splitLineAt(i, j int) {
	// Bounds check the input
	n := tg.length()
	if i < 0 || i >= n {
		err := fmt.Errorf("index %d out of bounds for grid of length %d", i, n)
		panic(err)
	}
	m := tg.widthAt(i)
	if j < 0 || j > m {
		err := fmt.Errorf("index %d out of bounds for row width %d at line %d", j, m, i)
		panic(err)
	}

	// Get the line & split
	old := tg.data[i]
	left := make([]rune, j)
	right := make([]rune, m-j)

	for idx := 0; idx < m; idx++ {
		if idx <= j {
			left[idx] = old[idx]
		} else {
			left[idx-j] = old[idx]
		}
	}

	// Add a new line after i
	tg.addLineAt(i + 1)

	// Set the new values
	tg.data[i] = left
	tg.data[j] = right
}

// setCharacterAt sets the value of the j-th character
// on the i-th line to the value c.
//
// Note: If i or j is out of bounds, setCharacterAt
// will panic.
func (tg *textGrid) setCharacterAt(i, j int, c rune) {
	// Bounds check the input
	n := tg.length()
	if i < 0 || i >= n {
		err := fmt.Errorf("index %d out of bounds for grid of length %d", i, n)
		panic(err)
	}
	m := tg.widthAt(i)
	if j < 0 || j >= m {
		err := fmt.Errorf("index %d out of bounds for row width %d at line %d", j, m, i)
		panic(err)
	}
	tg.data[i][j] = c
}
