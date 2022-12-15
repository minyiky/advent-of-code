package aocutils

type Vector struct {
	X, Y int
}

func NewVector(x, y int) Vector {
	return Vector{x, y}
}

func (v Vector) MDist(target Vector) int {
	x := Abs(v.X - target.X)
	y := Abs(v.Y - target.Y)
	return x + y
}
