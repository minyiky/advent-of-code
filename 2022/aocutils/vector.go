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

func (v Vector) Add(target Vector) Vector {
	x := v.X + target.X
	y := v.Y + target.Y
	return Vector{X: x, Y: y}
}

type Vector3D struct {
	X, Y, Z int
}

func NewVector3D(x, y, z int) Vector3D {
	return Vector3D{x, y, z}
}

func (v Vector3D) MDist(target Vector3D) int {
	x := Abs(v.X - target.X)
	y := Abs(v.Y - target.Y)
	z := Abs(v.Z - target.Z)
	return x + y + z
}

func (v Vector3D) Add(target Vector3D) Vector3D {
	x := v.X + target.X
	y := v.Y + target.Y
	z := v.Z + target.Z
	return Vector3D{X: x, Y: y, Z: z}
}
