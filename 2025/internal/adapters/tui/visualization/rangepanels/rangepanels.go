package rangepanels

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// ANSI color codes for coloring repeating digits
var digitColors = []string{
	"\033[91m", // Bright Red
	"\033[92m", // Bright Green
	"\033[93m", // Bright Yellow
	"\033[94m", // Bright Blue
	"\033[95m", // Bright Magenta
	"\033[96m", // Bright Cyan
	"\033[97m", // Bright White
}

const colorReset = "\033[0m"

type RangePanels struct {
	// No static state needed for range panels - all dynamic based on data
}

func New() *RangePanels {
	return &RangePanels{}
}

// colorizeRepeatingDigits adds color to repeating patterns
// For "123123 (×2)", the first "123" is one color, the second "123" is another color
// For "1111 (×4)", each "1" gets a different color
func colorizeRepeatingDigits(value string) (string, int) {
	// Parse the value to extract number and repetition count
	// Format: "number (×count)"
	parts := strings.Split(value, " (×")
	if len(parts) != 2 {
		return value, len(value) // Return as-is if format doesn't match
	}

	numStr := parts[0]
	repCountStr := strings.TrimSuffix(parts[1], ")")
	repCount, err := strconv.Atoi(repCountStr)
	if err != nil || repCount < 2 {
		return value, len(value) // Not a repeating pattern
	}

	// Calculate the base pattern length
	numLen := len(numStr)
	if numLen%repCount != 0 {
		return value, len(value) // Can't evenly divide
	}

	patternLen := numLen / repCount

	// Build colored string - each repetition group gets the same color
	var colored strings.Builder
	for i, ch := range numStr {
		// Which repetition are we in? (0, 1, 2, ...)
		repetitionIdx := i / patternLen
		colorIdx := repetitionIdx % len(digitColors)
		colored.WriteString(digitColors[colorIdx])
		colored.WriteRune(ch)
		colored.WriteString(colorReset)
	}

	// Add the repetition count part (uncolored)
	colored.WriteString(" (×")
	colored.WriteString(repCountStr)
	colored.WriteString(")")

	// Return colorized string and visible length (without ANSI codes)
	return colored.String(), len(value)
}

// Render renders a range-based visualization
func (rp *RangePanels) Render(currentRange string, rangeIndex, totalRanges int, panels map[string][]string, currentSum int) string {
	var sb strings.Builder

	// Show current range being processed with progress
	if currentRange == "" {
		currentRange = "None"
	}

	// Show progress through ranges
	if totalRanges > 0 {
		sb.WriteString(fmt.Sprintf("Range %d of %d: %s\n", rangeIndex+1, totalRanges, currentRange))
	} else {
		sb.WriteString(fmt.Sprintf("Current Range: %s\n", currentRange))
	}
	sb.WriteString(strings.Repeat("─", 70))
	sb.WriteString("\n\n")

	if len(panels) == 0 {
		sb.WriteString("No values found yet...\n\n")
		sb.WriteString(strings.Repeat("═", 70))
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("Current Sum: %d\n", currentSum))
		return sb.String()
	}

	// Extract and sort power ranges from panels
	powerRanges := make([]string, 0, len(panels))
	for powerRange := range panels {
		powerRanges = append(powerRanges, powerRange)
	}

	// Sort power ranges by their starting value
	sort.Slice(powerRanges, func(i, j int) bool {
		var startI, startJ int
		fmt.Sscanf(powerRanges[i], "%d-", &startI)
		fmt.Sscanf(powerRanges[j], "%d-", &startJ)
		return startI < startJ
	})

	// Determine panel width based on max value length
	const maxValuesPerPanel = 20 // Limit display to 20 values per panel
	const panelSpacing = 2

	// Calculate column widths for each panel (using visible length, not including ANSI codes)
	colWidths := make([]int, len(powerRanges))
	for i, powerRange := range powerRanges {
		// Min width is the header width
		colWidths[i] = len(powerRange) + 4 // "┌─ XX ─┐" format

		// Check values width (use visible length)
		values := panels[powerRange]
		for j, val := range values {
			if j >= maxValuesPerPanel {
				break
			}
			// Get visible length (without ANSI codes)
			_, visibleLen := colorizeRepeatingDigits(val)
			if visibleLen+4 > colWidths[i] { // +4 for "│ XX │"
				colWidths[i] = visibleLen + 4
			}
		}
		// Ensure minimum width
		if colWidths[i] < 10 {
			colWidths[i] = 10
		}
	}

	// Find max rows needed
	maxRows := 0
	for _, values := range panels {
		displayCount := len(values)
		if displayCount > maxValuesPerPanel {
			displayCount = maxValuesPerPanel + 1 // +1 for "..."
		}
		if displayCount > maxRows {
			maxRows = displayCount
		}
	}

	// Render headers (top borders with range labels)
	for i, powerRange := range powerRanges {
		headerText := fmt.Sprintf(" %s ", powerRange)
		headerWidth := colWidths[i] - 2 // -2 for "┌┐"
		padding := (headerWidth - len(headerText)) / 2

		sb.WriteString("┌")
		sb.WriteString(strings.Repeat("─", padding))
		sb.WriteString(headerText)
		sb.WriteString(strings.Repeat("─", headerWidth-padding-len(headerText)))
		sb.WriteString("┐")

		if i < len(powerRanges)-1 {
			sb.WriteString(strings.Repeat(" ", panelSpacing))
		}
	}
	sb.WriteString("\n")

	// Render value rows
	for row := 0; row < maxRows; row++ {
		for i, powerRange := range powerRanges {
			values := panels[powerRange]
			var cellContent string
			var visibleLen int

			if row < len(values) && row < maxValuesPerPanel {
				// Colorize the cell content
				cellContent, visibleLen = colorizeRepeatingDigits(values[row])
			} else if row == maxValuesPerPanel && len(values) > maxValuesPerPanel {
				cellContent = "..."
				visibleLen = 3
			} else {
				cellContent = ""
				visibleLen = 0
			}

			// Render cell
			sb.WriteString("│ ")
			sb.WriteString(cellContent)
			padding := colWidths[i] - visibleLen - 3 // -3 for "│ │", use visible length
			if padding > 0 {
				sb.WriteString(strings.Repeat(" ", padding))
			}
			sb.WriteString("│")

			if i < len(powerRanges)-1 {
				sb.WriteString(strings.Repeat(" ", panelSpacing))
			}
		}
		sb.WriteString("\n")
	}

	// Render bottom borders
	for i := range powerRanges {
		sb.WriteString("└")
		sb.WriteString(strings.Repeat("─", colWidths[i]-2))
		sb.WriteString("┘")

		if i < len(powerRanges)-1 {
			sb.WriteString(strings.Repeat(" ", panelSpacing))
		}
	}
	sb.WriteString("\n")

	// Show value counts
	sb.WriteString("\n")
	for i, powerRange := range powerRanges {
		count := len(panels[powerRange])
		countText := fmt.Sprintf("(%d)", count)
		padding := (colWidths[i] - len(countText)) / 2
		sb.WriteString(strings.Repeat(" ", padding))
		sb.WriteString(countText)

		if i < len(powerRanges)-1 {
			sb.WriteString(strings.Repeat(" ", colWidths[i]-padding-len(countText)+panelSpacing))
		}
	}
	sb.WriteString("\n\n")

	// Show current sum at bottom
	sb.WriteString(strings.Repeat("═", 70))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Current Sum: %d\n", currentSum))

	return sb.String()
}
