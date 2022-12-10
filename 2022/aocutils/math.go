package aocutils

import "golang.org/x/exp/constraints"

// Abs extends the math.Abs function with  generics
func Abs[T constraints.Float | constraints.Signed](v T) T {
	if v < 0 {
		v *= -1
	}
	return v
}
