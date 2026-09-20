package cmds

import (
	"fmt"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
)

func Find(state *EditorState, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: f <text>")
	}

	matches := 0
	text := strings.TrimSpace(strings.Join(args, " "))
	for i, line := range(state.Buffer) {
		if strings.Contains(line, text) {
			color.PrintGreenln(fmt.Sprintf("Match in line %d", i + 1))
			helpers.PrintFileLine(i + 1, line)
			matches++
		}
	}

	switch matches {
	case 0:
		color.PrintYellowln("No Matches!")
	case 1:
		color.PrintGreenln("1 Match")
	default:
		color.PrintGreenln(fmt.Sprintf("%d Matches.", matches))
	}

	return nil
}
