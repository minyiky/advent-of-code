package visualization

import (
	"fmt"
	"sort"
	"strings"
)

// RenderRangePanels renders a range-based visualization for Day 2
// showing the current range at top and panels organized by power ranges (side by side)
func RenderRangePanels(currentRange string, rangeIndex, totalRanges int, panels map[string][]string, currentSum int) string {
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

	// Calculate column widths for each panel
	colWidths := make([]int, len(powerRanges))
	for i, powerRange := range powerRanges {
		// Min width is the header width
		colWidths[i] = len(powerRange) + 4 // "┌─ XX ─┐" format

		// Check values width
		values := panels[powerRange]
		for j, val := range values {
			if j >= maxValuesPerPanel {
				break
			}
			if len(val)+4 > colWidths[i] { // +4 for "│ XX │"
				colWidths[i] = len(val) + 4
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

			if row < len(values) && row < maxValuesPerPanel {
				cellContent = values[row]
			} else if row == maxValuesPerPanel && len(values) > maxValuesPerPanel {
				cellContent = "..."
			} else {
				cellContent = ""
			}

			// Render cell
			sb.WriteString("│ ")
			sb.WriteString(cellContent)
			padding := colWidths[i] - len(cellContent) - 3 // -3 for "│ │"
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
