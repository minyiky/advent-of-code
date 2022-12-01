package dayone

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string
var lines []string

func init() {
	input = strings.ReplaceAll(input, "\r", "")
	lines = strings.Split(input, "\n")
}
