package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	rows    [][]int
	columns [][]int
	height  int
	width   int
}

func New(in string) (*Matrix, error) {
	matrix := Matrix{}
	lines := strings.Split(in, "\n")
	matrix.height = len(lines)
	matrix.width = len(strings.Split(strings.TrimSpace(lines[0]), " "))

	matrix.initilializeMatrix()

	for row, line := range lines {
		nums := strings.Split(strings.TrimSpace(line), " ")

		if len(nums) != matrix.width {
			return nil, errors.New("uneven rows")
		}

		for col, n := range nums {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			matrix.addToMatrix(num, col, row)
		}
	}

	return &matrix, nil
}

func (m *Matrix) Rows() [][]int {
	var response [][]int

	for _, row := range m.rows {
		tmp := make([]int, m.width)
		copy(tmp, row)
		response = append(response, tmp)
	}
	return response
}

func (m *Matrix) Cols() [][]int {
	var response [][]int

	for _, col := range m.columns {
		tmp := make([]int, m.height)
		copy(tmp, col)
		response = append(response, tmp)
	}

	return response
}

func (m *Matrix) Set(row, col, value int) bool {
	if m.height-1 < row || m.width-1 < col || row < 0 || col < 0 {
		return false
	}

	m.rows[row][col] = value
	m.columns[col][row] = value
	return true
}

func (m *Matrix) initilializeMatrix() {
	a := make([][]int, m.height)
	for i := range a {
		a[i] = make([]int, m.width)
	}

	m.rows = a

	b := make([][]int, m.width)
	for i := range b {
		b[i] = make([]int, m.height)
	}

	m.columns = b
}

func (m *Matrix) addToMatrix(num, col, row int) {
	m.rows[row][col] = num
	m.columns[col][row] = num
}
