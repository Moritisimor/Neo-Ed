package tokenizing

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/cmds"
)

func Dispatch(tokens []string) {
	switch tokens[0] {
	default:
		color.PrintRedln("Unknown command: " + tokens[0])

	case "help":
		cmds.PrintHelp()
	}
}
