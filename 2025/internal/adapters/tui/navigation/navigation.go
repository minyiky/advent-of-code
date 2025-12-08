package navigation

import (
	"fmt"
	"strings"
)

type Control struct {
	Keys        []string
	Instruction string
}

func (c *Control) String() string {
	return fmt.Sprintf("%s to %s", strings.Join(c.Keys, "/"), c.Instruction)
}

func DefaultControls() []Control {
	return []Control{
		{
			Keys:        []string{"ESC"},
			Instruction: "go back",
		}, {
			Keys:        []string{"q"},
			Instruction: "quit",
		},
	}
}
