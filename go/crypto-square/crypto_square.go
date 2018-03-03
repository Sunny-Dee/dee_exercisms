package cryptosquare

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type matrix struct {
	lines   [][]byte
	rows    int
	colSize int
}

// Encode implements the square code
// method for composing secret messages.
func Encode(plain string) string {
	if plain == "" {
		return ""
	}

	plain = strings.ToLower(plain)

	r := regexp.MustCompile("[[:punct:]\\s]*")
	clean := r.ReplaceAllString(plain, "")

	rows, columns := calcRowsCols(clean)
	mtx := newMatrix(rows, columns)
	mtx.populate(clean)
	return mtx.String()
}

func calcRowsCols(clean string) (int, int) {
	l := len(clean)
	sqrt := math.Sqrt(float64(l))
	row := int(sqrt)
	if sqrt == float64(int(sqrt)) {
		return row, row
	}

	for {
		if row*row >= l {
			return row, row
		} else if (row+1)*row >= l {
			return row + 1, row
		} else {
			row++
		}
	}
}

func newMatrix(rows, cols int) matrix {
	lines := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		lines[i] = make([]byte, cols)
	}

	return matrix{
		lines:   lines,
		rows:    rows,
		colSize: cols,
	}
}

func (m *matrix) populate(str string) {
	col := 0
	row := 0
	for i := 0; i < len(str); i++ {
		m.lines[row][col] = str[i]
		row++

		if row%m.rows == 0 {
			col++
			row = 0
		}
	}
}

func (m *matrix) String() string {
	var sb strings.Builder

	for i := 0; i < len(m.lines)-1; i++ {
		line := bytes.Trim(m.lines[i], "\x00")
		sb.WriteString(fmt.Sprintf("%s ", line))
	}

	line := bytes.Trim(m.lines[len(m.lines)-1], "\x00")
	sb.WriteString(fmt.Sprintf("%s", line))

	return sb.String()
}
