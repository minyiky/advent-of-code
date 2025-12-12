package day10

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	results := make(chan int, len(lines))

	wg := sync.WaitGroup{}

	for _, line := range lines {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			parts := strings.Split(line, " ")

			target := make([]int, 0)
			jolts := strings.Split(strings.Trim(parts[len(parts)-1], "{}"), ",")

			for _, jolt := range jolts {
				num, err := strconv.Atoi(jolt)
				if err != nil {
					return
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

			result, ok := solveBranchAndBound(instructions, len(target), target)
			if ok {
				results <- result
			}
		}(line)
	}

	wg.Wait()
	close(results)

	for r := range results {
		value += r
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
	pivotCol := make([]int, numBulbs)
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

		// Swap rows (only if needed)
		if maxRow != pivotRow {
			matrix[pivotRow], matrix[maxRow] = matrix[maxRow], matrix[pivotRow]
		}

		// Scale pivot row to make pivot = 1
		scale := matrix[pivotRow][col]
		for j := col; j <= numInstr; j++ { // Optimization: start from col
			matrix[pivotRow][j] /= scale
		}

		// Eliminate column in all other rows
		for row := 0; row < numBulbs; row++ {
			if row != pivotRow && math.Abs(matrix[row][col]) > 1e-10 {
				factor := matrix[row][col]
				for j := col; j <= numInstr; j++ { // Optimization: start from col
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

	maxTarget := 0
	for _, t := range target {
		if t > maxTarget {
			maxTarget = t
		}
	}

	maxFreeVal := maxTarget

	bestSum := -1

	// Try to find an initial solution with all free vars = 0
	{
		solution := make([]float64, numInstr)
		for row := pivotRow - 1; row >= 0; row-- {
			col := pivotCol[row]
			if col >= 0 {
				solution[col] = matrix[row][numInstr]
			}
		}
		valid := true
		sum := 0
		for _, v := range solution {
			rounded := int(math.Round(v))
			if rounded < 0 || math.Abs(v-float64(rounded)) > 1e-6 {
				valid = false
				break
			}
			sum += rounded
		}
		if valid {
			bestSum = sum
		}
	}

	// Reusable solution slice
	solution := make([]float64, numInstr)

	var search func(freeIdx int, freeVals []int, partialSum int)
	search = func(freeIdx int, freeVals []int, partialSum int) {
		// Prune: if partial sum already exceeds best, skip
		if bestSum >= 0 && partialSum >= bestSum {
			return
		}

		if freeIdx == len(freeVars) {
			// Compute solution with these free variable values
			for i := 0; i < numInstr; i++ {
				solution[i] = 0
			}
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
			// Prune: if adding this v already exceeds best, skip rest
			if bestSum >= 0 && partialSum+v >= bestSum {
				break
			}
			freeVals[freeIdx] = v
			search(freeIdx+1, freeVals, partialSum+v)
		}
	}

	if len(freeVars) == 0 {
		// Unique solution - just extract it
		for i := 0; i < numInstr; i++ {
			solution[i] = 0
		}
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

	search(0, make([]int, len(freeVars)), 0)
	if bestSum < 0 {
		return 0, false
	}
	return bestSum, true
}

// Branch and Bound with LP Relaxation solver
func solveBranchAndBound(instructions [][]int, numBulbs int, target []int) (int, bool) {
	numVars := len(instructions)

	// Build constraint matrix A and vector b for Ax = b
	A := make([][]float64, numBulbs)
	for i := range A {
		A[i] = make([]float64, numVars)
	}
	b := make([]float64, numBulbs)
	for i, t := range target {
		b[i] = float64(t)
	}
	for instrIdx, instr := range instructions {
		for _, bulb := range instr {
			A[bulb][instrIdx] = 1
		}
	}

	// Objective: minimize sum of all variables (coefficients all 1)
	c := make([]float64, numVars)
	for i := range c {
		c[i] = 1
	}

	bestIntegerSum := -1

	// Branch and bound state
	type Node struct {
		lowerBounds []float64 // x_i >= lowerBounds[i]
		upperBounds []float64 // x_i <= upperBounds[i]
	}

	// Initialize with no extra bounds
	maxBound := float64(0)
	for _, t := range target {
		if float64(t) > maxBound {
			maxBound = float64(t)
		}
	}

	initialNode := Node{
		lowerBounds: make([]float64, numVars),
		upperBounds: make([]float64, numVars),
	}
	for i := range initialNode.upperBounds {
		initialNode.upperBounds[i] = maxBound * 2 // Upper bound
	}

	// Stack for DFS
	stack := []Node{initialNode}

	const intTol = 1e-6

	for len(stack) > 0 {
		// Pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Solve LP relaxation with current bounds
		solution, objVal, feasible := solveLP(A, b, c, node.lowerBounds, node.upperBounds)
		if !feasible {
			continue // Infeasible, prune
		}

		// Prune if lower bound >= best integer solution
		if bestIntegerSum >= 0 && objVal >= float64(bestIntegerSum) {
			continue
		}

		// Clamp tiny negative values to zero for numerical stability
		for i, v := range solution {
			if v > -intTol && v < 0 {
				solution[i] = 0
			}
		}

		// Check if solution is integer and find best fractional variable to branch on
		allInteger := true
		fractionalIdx := -1
		bestFrac := 0.0

		for i, v := range solution {
			// Truly negative values indicate infeasibility for this branch
			if v < -intTol {
				allInteger = false
				// Still need to find a variable to branch on
				if fractionalIdx < 0 {
					fractionalIdx = i
					bestFrac = 0.5 // Force branching on this
				}
				continue
			}

			frac := math.Abs(v - math.Round(v))
			if frac > intTol {
				allInteger = false
				// Pick the most fractional variable (furthest from integer)
				if frac > bestFrac {
					bestFrac = frac
					fractionalIdx = i
				}
			}
		}

		if allInteger {
			// Found integer solution - verify it satisfies Ax = b
			sum := 0
			valid := true
			intSol := make([]int, len(solution))
			for i, v := range solution {
				if v < -intTol {
					valid = false
					break
				}
				intSol[i] = int(math.Round(v))
				sum += intSol[i]
			}

			// Verify constraints
			if valid {
				for i := 0; i < numBulbs; i++ {
					lhs := 0.0
					for j := 0; j < numVars; j++ {
						lhs += A[i][j] * float64(intSol[j])
					}
					if math.Abs(lhs-b[i]) > intTol {
						valid = false
						break
					}
				}
			}

			if valid && (bestIntegerSum < 0 || sum < bestIntegerSum) {
				bestIntegerSum = sum
			}
			continue
		}

		if fractionalIdx < 0 {
			// Should not happen after the fixes above, but guard anyway
			continue
		}

		// Branch on fractional variable
		fracVal := solution[fractionalIdx]
		floorVal := math.Floor(fracVal)
		ceilVal := math.Ceil(fracVal)

		// Branch 1: x_i <= floor
		if floorVal >= node.lowerBounds[fractionalIdx] {
			newNode := Node{
				lowerBounds: make([]float64, numVars),
				upperBounds: make([]float64, numVars),
			}
			copy(newNode.lowerBounds, node.lowerBounds)
			copy(newNode.upperBounds, node.upperBounds)
			newNode.upperBounds[fractionalIdx] = floorVal
			stack = append(stack, newNode)
		}

		// Branch 2: x_i >= ceil
		if ceilVal <= node.upperBounds[fractionalIdx] {
			newNode := Node{
				lowerBounds: make([]float64, numVars),
				upperBounds: make([]float64, numVars),
			}
			copy(newNode.lowerBounds, node.lowerBounds)
			copy(newNode.upperBounds, node.upperBounds)
			newNode.lowerBounds[fractionalIdx] = ceilVal
			stack = append(stack, newNode)
		}
	}

	if bestIntegerSum < 0 {
		return 0, false
	}
	return bestIntegerSum, true
}

// Simplex solver: minimize c'x subject to Ax = b, lb <= x <= ub
// Converts bounded problem to standard form and uses two-phase simplex
func solveLP(A [][]float64, b []float64, c []float64, lb []float64, ub []float64) ([]float64, float64, bool) {
	m := len(A) // number of original constraints
	n := len(c) // number of original variables

	if m == 0 || n == 0 {
		return nil, 0, false
	}

	const eps = 1e-9
	const maxIter = 10000

	// Shift variables: y = x - lb, so y >= 0 and y <= ub - lb
	// Original constraints: Ay = b - A*lb
	bShifted := make([]float64, m)
	for i := 0; i < m; i++ {
		bShifted[i] = b[i]
		for j := 0; j < n; j++ {
			bShifted[i] -= A[i][j] * lb[j]
		}
	}

	ubShifted := make([]float64, n)
	for j := 0; j < n; j++ {
		ubShifted[j] = ub[j] - lb[j]
		if ubShifted[j] < -eps {
			return nil, 0, false // Infeasible bounds
		}
	}

	// Convert upper bounds to equality constraints with slack variables
	// For each j: y_j + s_j = ubShifted[j], where s_j >= 0
	// New problem has n original vars + n slack vars = 2n vars
	// And m original constraints + n bound constraints = m + n constraints
	
	mNew := m + n  // total constraints
	nNew := 2 * n  // original vars + slack vars

	// Build new constraint matrix
	ANew := make([][]float64, mNew)
	for i := range ANew {
		ANew[i] = make([]float64, nNew)
	}
	bNew := make([]float64, mNew)

	// Original constraints: A*y = bShifted (y are first n vars)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ANew[i][j] = A[i][j]
		}
		bNew[i] = bShifted[i]
	}

	// Bound constraints: y_j + s_j = ubShifted[j]
	for j := 0; j < n; j++ {
		ANew[m+j][j] = 1       // y_j
		ANew[m+j][n+j] = 1     // s_j
		bNew[m+j] = ubShifted[j]
	}

	// Objective: minimize c'y (slack vars have 0 cost)
	cNew := make([]float64, nNew)
	for j := 0; j < n; j++ {
		cNew[j] = c[j]
	}

	// Now solve standard LP: min cNew'z s.t. ANew*z = bNew, z >= 0
	// Using two-phase simplex

	// Phase 1: Find initial BFS using artificial variables
	numCols := nNew + mNew + 1 // vars + artificials + RHS
	tableau := make([][]float64, mNew+1)
	for i := range tableau {
		tableau[i] = make([]float64, numCols)
	}

	// Fill constraint rows
	for i := 0; i < mNew; i++ {
		sign := 1.0
		if bNew[i] < 0 {
			sign = -1.0
		}
		for j := 0; j < nNew; j++ {
			tableau[i][j] = sign * ANew[i][j]
		}
		tableau[i][nNew+i] = 1.0 // Artificial variable
		tableau[i][numCols-1] = sign * bNew[i]
	}

	// Phase 1 objective: minimize sum of artificials
	for i := 0; i < mNew; i++ {
		for j := 0; j < numCols; j++ {
			tableau[mNew][j] -= tableau[i][j]
		}
	}
	for i := 0; i < mNew; i++ {
		tableau[mNew][nNew+i] = 0
	}

	// Basis tracking
	basis := make([]int, mNew)
	for i := 0; i < mNew; i++ {
		basis[i] = nNew + i
	}

	// Phase 1 simplex iterations
	phase1Optimal := false
	for iter := 0; iter < maxIter; iter++ {
		// Find entering variable (most negative reduced cost)
		enterCol := -1
		minCoeff := -eps
		for j := 0; j < nNew+mNew; j++ {
			if tableau[mNew][j] < minCoeff {
				minCoeff = tableau[mNew][j]
				enterCol = j
			}
		}
		if enterCol < 0 {
			phase1Optimal = true
			break // Optimal for Phase 1
		}

		// Find leaving variable (minimum ratio test)
		leaveRow := -1
		minRatio := math.MaxFloat64
		for i := 0; i < mNew; i++ {
			if tableau[i][enterCol] > eps {
				ratio := tableau[i][numCols-1] / tableau[i][enterCol]
				if ratio >= 0 && ratio < minRatio {
					minRatio = ratio
					leaveRow = i
				}
			}
		}
		if leaveRow < 0 {
			return nil, 0, false // Unbounded in Phase 1
		}

		// Pivot
		pivot := tableau[leaveRow][enterCol]
		for j := 0; j < numCols; j++ {
			tableau[leaveRow][j] /= pivot
		}
		for i := 0; i <= mNew; i++ {
			if i != leaveRow {
				factor := tableau[i][enterCol]
				for j := 0; j < numCols; j++ {
					tableau[i][j] -= factor * tableau[leaveRow][j]
				}
			}
		}
		basis[leaveRow] = enterCol
	}
	if !phase1Optimal {
		return nil, 0, false // Hit iteration limit in Phase 1
	}

	// Check Phase 1 feasibility
	if tableau[mNew][numCols-1] > eps {
		return nil, 0, false // Infeasible
	}

	// Remove degenerate artificial variables from basis
	// If an artificial is in basis with value 0, pivot it out for a real variable
	for i := 0; i < mNew; i++ {
		if basis[i] >= nNew { // Artificial variable in basis
			// Try to find a non-artificial column to pivot in
			for j := 0; j < nNew; j++ {
				if math.Abs(tableau[i][j]) > eps {
					// Pivot this column into the basis
					pivot := tableau[i][j]
					for k := 0; k < numCols; k++ {
						tableau[i][k] /= pivot
					}
					for row := 0; row <= mNew; row++ {
						if row != i {
							factor := tableau[row][j]
							for k := 0; k < numCols; k++ {
								tableau[row][k] -= factor * tableau[i][k]
							}
						}
					}
					basis[i] = j
					break
				}
			}
		}
	}

	// Phase 2: Optimize original objective
	for j := 0; j < numCols; j++ {
		tableau[mNew][j] = 0
	}
	for j := 0; j < nNew; j++ {
		tableau[mNew][j] = cNew[j]
	}

	// Eliminate basic variables from objective
	for i := 0; i < mNew; i++ {
		col := basis[i]
		if col < nNew {
			factor := tableau[mNew][col]
			for j := 0; j < numCols; j++ {
				tableau[mNew][j] -= factor * tableau[i][j]
			}
		}
	}

	// Phase 2 simplex iterations
	phase2Optimal := false
	for iter := 0; iter < maxIter; iter++ {
		// Find entering variable
		enterCol := -1
		minCoeff := -eps
		for j := 0; j < nNew; j++ {
			if tableau[mNew][j] < minCoeff {
				minCoeff = tableau[mNew][j]
				enterCol = j
			}
		}
		if enterCol < 0 {
			phase2Optimal = true
			break // Optimal
		}

		// Find leaving variable
		leaveRow := -1
		minRatio := math.MaxFloat64
		for i := 0; i < mNew; i++ {
			if tableau[i][enterCol] > eps {
				ratio := tableau[i][numCols-1] / tableau[i][enterCol]
				if ratio >= 0 && ratio < minRatio {
					minRatio = ratio
					leaveRow = i
				}
			}
		}
		if leaveRow < 0 {
			return nil, 0, false // Unbounded
		}

		// Pivot
		pivot := tableau[leaveRow][enterCol]
		for j := 0; j < numCols; j++ {
			tableau[leaveRow][j] /= pivot
		}
		for i := 0; i <= mNew; i++ {
			if i != leaveRow {
				factor := tableau[i][enterCol]
				for j := 0; j < numCols; j++ {
					tableau[i][j] -= factor * tableau[leaveRow][j]
				}
			}
		}
		basis[leaveRow] = enterCol
	}
	if !phase2Optimal {
		return nil, 0, false // Hit iteration limit in Phase 2
	}

	// Verify dual optimality: no negative reduced costs
	for j := 0; j < nNew; j++ {
		if tableau[mNew][j] < -eps {
			return nil, 0, false // Not actually optimal
		}
	}

	// Extract solution for original shifted variables y
	y := make([]float64, n)
	for i := 0; i < mNew; i++ {
		col := basis[i]
		if col < n {
			y[col] = tableau[i][numCols-1]
		}
	}

	// Convert back to original coordinates: x = y + lb
	x := make([]float64, n)
	for j := 0; j < n; j++ {
		x[j] = y[j] + lb[j]
		// Verify bounds
		if x[j] < lb[j]-eps || x[j] > ub[j]+eps {
			return nil, 0, false
		}
		// Clamp to bounds
		if x[j] < lb[j] {
			x[j] = lb[j]
		}
		if x[j] > ub[j] {
			x[j] = ub[j]
		}
	}

	// Verify equality constraints: A x ≈ b (in original coordinates)
	for i := 0; i < m; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += A[i][j] * x[j]
		}
		if math.Abs(sum-b[i]) > 1e-6 {
			return nil, 0, false // LP solution doesn't satisfy constraints
		}
	}

	// Compute objective
	objVal := 0.0
	for j := 0; j < n; j++ {
		objVal += c[j] * x[j]
	}

	return x, objVal, true
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
