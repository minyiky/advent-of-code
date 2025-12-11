package day10

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		parts := strings.Split(line, " ")

		target := make([]int, 0)
		jolts := strings.Split(strings.Trim(parts[len(parts)-1], "{}"), ",")

		for _, jolt := range jolts {
			num, err := strconv.Atoi(jolt)
			if err != nil {
				return 0, err
			}
			target = append(target, num)
		}

		instructions := make([][]int, 0, len(parts)-2)
		for _, p := range parts[1 : len(parts)-1] {
			var instruction []int
			for _, c := range p {
				if c >= '0' && c <= '9' {
					instruction = append(instruction, int(c-'0'))
				}
			}
			if len(instruction) > 0 {
				instructions = append(instructions, instruction)
			}

		}

		if result, ok := solveLinearSystem(instructions, len(target), target); ok {
			value += result
		}
	}

	return value, nil
}

func solveLinearSystem(instructions [][]int, numBulbs int, target []int) (int, bool) {
	numInstr := len(instructions)

	// Build augmented matrix [A|b] (rows=bulbs, cols=instructions+1)
	matrix := make([][]float64, numBulbs)
	for i := range matrix {
		matrix[i] = make([]float64, numInstr+1)
		matrix[i][numInstr] = float64(target[i])
	}
	for instrIdx, instr := range instructions {
		for _, bulb := range instr {
			matrix[bulb][instrIdx] = 1
		}
	}

	// Track which columns have pivots and their pivot rows
	pivotCol := make([]int, numBulbs) // pivotCol[row] = column of pivot, or -1
	for i := range pivotCol {
		pivotCol[i] = -1
	}

	// Gauss-Jordan elimination
	pivotRow := 0
	for col := 0; col < numInstr && pivotRow < numBulbs; col++ {
		// Find pivot (largest absolute value in column)
		maxRow := pivotRow
		for row := pivotRow + 1; row < numBulbs; row++ {
			if math.Abs(matrix[row][col]) > math.Abs(matrix[maxRow][col]) {
				maxRow = row
			}
		}

		if math.Abs(matrix[maxRow][col]) < 1e-10 {
			continue // No pivot in this column (free variable)
		}

		// Swap rows
		matrix[pivotRow], matrix[maxRow] = matrix[maxRow], matrix[pivotRow]

		// Scale pivot row to make pivot = 1
		scale := matrix[pivotRow][col]
		for j := range matrix[pivotRow] {
			matrix[pivotRow][j] /= scale
		}

		// Eliminate column in all other rows
		for row := 0; row < numBulbs; row++ {
			if row != pivotRow && math.Abs(matrix[row][col]) > 1e-10 {
				factor := matrix[row][col]
				for j := range matrix[row] {
					matrix[row][j] -= factor * matrix[pivotRow][j]
				}
			}
		}

		pivotCol[pivotRow] = col
		pivotRow++
	}

	// Identify free variables (columns without pivots)
	hasPivot := make([]bool, numInstr)
	for row := 0; row < pivotRow; row++ {
		if pivotCol[row] >= 0 {
			hasPivot[pivotCol[row]] = true
		}
	}

	var freeVars []int
	for col := 0; col < numInstr; col++ {
		if !hasPivot[col] {
			freeVars = append(freeVars, col)
		}
	}

	// Search for minimum sum solution with non-negative integers
	// Try all combinations of free variables from 0 to max reasonable value
	maxFreeVal := 0
	for _, t := range target {
		if t > maxFreeVal {
			maxFreeVal = t
		}
	}

	bestSum := -1

	var search func(freeIdx int, freeVals []int)
	search = func(freeIdx int, freeVals []int) {
		if freeIdx == len(freeVars) {
			// Compute solution with these free variable values
			solution := make([]float64, numInstr)
			for i, col := range freeVars {
				solution[col] = float64(freeVals[i])
			}

			// Back-substitute to find pivot variables
			for row := pivotRow - 1; row >= 0; row-- {
				col := pivotCol[row]
				if col < 0 {
					continue
				}
				val := matrix[row][numInstr]
				for j := col + 1; j < numInstr; j++ {
					val -= matrix[row][j] * solution[j]
				}
				solution[col] = val
			}

			// Check if valid (non-negative integers)
			sum := 0
			for _, v := range solution {
				rounded := int(math.Round(v))
				if rounded < 0 || math.Abs(v-float64(rounded)) > 1e-6 {
					return // Invalid solution
				}
				sum += rounded
			}

			if bestSum < 0 || sum < bestSum {
				bestSum = sum
			}
			return
		}

		// Try values for this free variable
		for v := 0; v <= maxFreeVal; v++ {
			freeVals[freeIdx] = v
			search(freeIdx+1, freeVals)
		}
	}

	if len(freeVars) == 0 {
		// Unique solution - just extract it
		solution := make([]float64, numInstr)
		for row := 0; row < pivotRow; row++ {
			col := pivotCol[row]
			if col >= 0 {
				solution[col] = matrix[row][numInstr]
			}
		}

		sum := 0
		for _, v := range solution {
			rounded := int(math.Round(v))
			if rounded < 0 || math.Abs(v-float64(rounded)) > 1e-6 {
				return 0, false
			}
			sum += rounded
		}
		return sum, true
	}

	search(0, make([]int, len(freeVars)))
	if bestSum < 0 {
		return 0, false
	}
	return bestSum, true
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
