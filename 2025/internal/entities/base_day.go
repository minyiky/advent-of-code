package entities

import "time"

type day interface {
	Part1Val(input []string) (int, error)
	Part2Val(input []string) (int, error)
}

type visualizableDay interface {
	day
	Part1ValWithViz(input []string, vizChan chan<- VizFrame, speed float64) (int, error)
	Part2ValWithViz(input []string, vizChan chan<- VizFrame, speed float64) (int, error)
}

// baseDay provides common Day interface implementation using generics
type baseDay[T day] struct {
	impl T
	n    int
}

// NewBaseDay creates a new BaseDay with the given implementation
func NewBaseDay[T day](impl T, n int) *baseDay[T] {
	return &baseDay[T]{
		impl: impl,
		n:    n,
	}
}

func (d *baseDay[_]) Number() int {
	return d.n
}

func (d *baseDay[_]) RunPart1(input []string) (int, error) {
	return d.impl.Part1Val(input)
}

func (d *baseDay[_]) RunPart2(input []string) (int, error) {
	return d.impl.Part2Val(input)
}

func (d *baseDay[_]) TimeAndRunPart(input []string, part int) (time.Duration, int, error) {
	start := time.Now()
	var result int
	var err error

	if part == 1 {
		result, err = d.impl.Part1Val(input)
	} else {
		result, err = d.impl.Part2Val(input)
	}

	duration := time.Since(start)
	return duration, result, err
}

func (d *baseDay[_]) RunPartWithVisualization(input []string, part int, vizChan chan<- VizFrame, speed float64) (time.Duration, int, error) {
	vizImpl, ok := any(d.impl).(visualizableDay)
	if !ok {
		return d.TimeAndRunPart(input, part)
	}

	start := time.Now()
	result, err := func() (int, error) {
		if part == 1 {
			return vizImpl.Part1ValWithViz(input, vizChan, speed)
		}
		return vizImpl.Part2ValWithViz(input, vizChan, speed)
	}()

	duration := time.Since(start)
	return duration, result, err

}
