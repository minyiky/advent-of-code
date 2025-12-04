package day04

var (
	Up = Point{
		X: 0,
		Y: -1,
	}
	Down = Point{
		X: 0,
		Y: 1,
	}
	Left = Point{
		X: -1,
		Y: 0,
	}
	Right = Point{
		X: 1,
		Y: 0,
	}
)

var (
	UpLeft = Point{
		X: -1,
		Y: -1,
	}
	UpRight = Point{
		X: 1,
		Y: -1,
	}
	DownLeft = Point{
		X: -1,
		Y: 1,
	}
	DownRight = Point{
		X: 1,
		Y: 1,
	}
)

var Cardinals = []Point{Up, Down, Left, Right}
var Diagonals = []Point{UpLeft, UpRight, DownLeft, DownRight}
var AllDirs = append(Cardinals, Diagonals...)

type Point struct {
	X int
	Y int
}

func (p *Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}
