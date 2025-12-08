package visualization

type Canvas [][]rune

func (c Canvas) Copy() Canvas {
	newC := make(Canvas, len(c))
	rowLen := len(c[0])
	for i := range c {
		newC[i] = make([]rune, rowLen)
		copy(newC[i], c[i])
	}
	return newC
}

type CanvasFactory struct{}

func NewCanvasFactory() *CanvasFactory {
	return &CanvasFactory{}
}

func (f *CanvasFactory) New(w, h int) Canvas {
	// Create a 2D canvas
	canvas := make(Canvas, h)
	for i := range canvas {
		canvas[i] = make([]rune, w)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}

	return canvas
}
