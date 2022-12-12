package aocutils

type Vector struct {
	X, Y int
}

func NewVector(x, y int) Vector {
	return Vector{x, y}
}
