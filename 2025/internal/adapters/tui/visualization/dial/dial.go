package dial

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/visualization"
)

const (
	radiusY = 12
	radiusX = 24 // Wider to compensate for terminal character aspect ratio
	centerX = radiusX
	centerY = radiusY
	width   = radiusX*2 + 1
	height  = radiusY*2 + 1

	// TODO make input
	steps = 100
)

type CanvasFactory interface {
	New(w, h int) visualization.Canvas
}

type Dial struct {
	baseCanvas visualization.Canvas
}

func New() *Dial {
	cf := visualization.NewCanvasFactory()

	baseCanvas := cf.New(width, height)

	// Draw the circle
	for angle := 0.0; angle < 360; angle += 1 {
		rad := angle * math.Pi / 180
		x := int(math.Round(float64(centerX) + float64(radiusX)*math.Cos(rad)))
		y := int(math.Round(float64(centerY) + float64(radiusY)*math.Sin(rad)))
		if x >= 0 && x < width && y >= 0 && y < height {
			baseCanvas[y][x] = '·'
		}
	}

	markers := [4]string{
		"0",
		strconv.Itoa(steps / 4),
		strconv.Itoa(2 * steps / 4),
		strconv.Itoa(3 * steps / 4),
	}

	for i, marker := range markers {
		angle := float64((i - 1) * 90)
		rad := angle * math.Pi / 180

		markerX := int(math.Round(float64(centerX) + float64(radiusX)*math.Cos(rad)))
		markerY := int(math.Round(float64(centerY) + float64(radiusY)*math.Sin(rad)))

		// Draw the label, centering multi-digit numbers
		for i, ch := range marker {
			offsetX := markerX + i - len(marker)/2
			if offsetX >= 0 && offsetX < width && markerY >= 0 && markerY < height {
				baseCanvas[markerY][offsetX] = ch
			}
		}
	}

	// Draw center point
	baseCanvas[centerY][centerX] = '+'

	return &Dial{
		baseCanvas: baseCanvas,
	}
}

// RenderDial renders a circular dial visualization with an instruction panel
// position is the current animated position, finalPosition is what to display at bottom
func (d *Dial) Render(position, counter, maxValue, finalPosition int, instructions []string, currentInstrIdx int) string {
	var sb strings.Builder

	// Draw position markers at key positions (0, 25, 50, 75)
	// Place them slightly outside the circle to avoid cutoff
	canvas := d.baseCanvas.Copy()

	// Draw the clock hand showing current position
	// Position 0 is at top (-90 degrees), advances clockwise
	handAngle := (float64(position)/float64(maxValue))*360 - 90
	handAngleRad := handAngle * math.Pi / 180
	handLengthX := float64(radiusX) - 2
	handLengthY := float64(radiusY) - 2

	// Draw smooth hand line from center outward
	for i := 0.0; i <= float64(radiusY); i += 0.5 {
		t := i / float64(radiusY) // Normalized distance from center (0 to 1)
		x := int(math.Round(float64(centerX) + handLengthX*t*math.Cos(handAngleRad)))
		y := int(math.Round(float64(centerY) + handLengthY*t*math.Sin(handAngleRad)))
		if x >= 0 && x < width && y >= 0 && y < height {
			canvas[y][x] = '█'
		}
	}

	// Draw counter in center (3 lines for vertical centering)
	counterStr := fmt.Sprintf("%d", counter)
	counterLen := len(counterStr)
	startX := centerX - counterLen/2

	if centerY >= 1 && centerY < height-1 {
		// Clear space for counter
		for x := max(0, startX-1); x < min(width, startX+counterLen+1); x++ {
			canvas[centerY-1][x] = ' '
			canvas[centerY][x] = ' '
			canvas[centerY+1][x] = ' '
		}

		// Draw counter
		for i, ch := range counterStr {
			x := startX + i
			if x >= 0 && x < width {
				canvas[centerY][x] = ch
			}
		}
	}

	// Convert canvas to string with instruction panel on the right
	dialLines := make([]string, len(canvas))
	for i, row := range canvas {
		dialLines[i] = string(row)
	}

	// Create instruction panel
	const panelLines = 5
	const contextBefore = 2
	const contextAfter = 2

	// Build instruction panel
	instrPanel := make([]string, height)
	if len(instructions) > 0 && currentInstrIdx >= 0 {
		// Calculate which instructions to show
		startIdx := max(0, currentInstrIdx-contextBefore)
		endIdx := min(len(instructions), currentInstrIdx+contextAfter+1)

		// Center the panel vertically
		panelStartLine := (height - panelLines) / 2

		lineIdx := 0
		for i := startIdx; i < endIdx && lineIdx < panelLines; i++ {
			var instrLine string
			if i == currentInstrIdx {
				// Highlight current instruction with bold, bright colors, and arrows
				instrLine = fmt.Sprintf("  \033[1;33m▶\033[0m \033[1;97m%-6s\033[0m \033[1;33m◀\033[0m", instructions[i])
			} else {
				// Dimmed past/future instructions
				instrLine = fmt.Sprintf("    \033[2m%-6s\033[0m  ", instructions[i])
			}

			if panelStartLine+lineIdx >= 0 && panelStartLine+lineIdx < height {
				instrPanel[panelStartLine+lineIdx] = instrLine
			}
			lineIdx++
		}
	}

	// Combine dial and instruction panel
	for i := 0; i < len(dialLines); i++ {
		sb.WriteString(dialLines[i])
		if len(instrPanel[i]) > 0 {
			sb.WriteString("  ")
			sb.WriteString(instrPanel[i])
		}
		if i < len(dialLines)-1 {
			sb.WriteRune('\n')
		}
	}

	// Add position info below (show final position for the current move)
	sb.WriteString(fmt.Sprintf("\n\nPosition: %d  Counter: %d", finalPosition, counter))

	return sb.String()
}
