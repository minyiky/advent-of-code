package day04

import (
	"strconv"
	"strings"
)

type cell struct {
	row, col int
	marked   bool
}

type board struct {
	cells   map[string]cell
	columns []int
	rows    []int

	numToWin int
	complete bool
}

func newBoard(lines []string) (*board, error) {
	cells := make(map[string]cell)

	for i, line := range lines {
		vals := strings.Fields(line)
		for j, val := range vals {
			cells[val] = cell{
				row:    i,
				col:    j,
				marked: false,
			}
		}
	}

	b := &board{
		cells:    cells,
		columns:  make([]int, len(lines)),
		rows:     make([]int, len(lines)),
		numToWin: len(lines),
	}

	return b, nil
}

func (b *board) mark(call string) {
	cell, ok := b.cells[call]
	if ok {
		cell.marked = true
		b.cells[call] = cell
		b.rows[cell.row]++
		b.columns[cell.col]++
		if b.columns[cell.col] == b.numToWin || b.rows[cell.row] == b.numToWin {
			b.complete = true
		}
	}
}

func (b *board) score() int {
	result := 0
	for k, v := range b.cells {
		if !v.marked {
			val, _ := strconv.Atoi(k)
			result += val
		}

	}
	return result
}
