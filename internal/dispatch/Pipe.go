package dispatch

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/cmds"
)

func Pipe(buf *[]string, args []string) {
	switch args[0] {
	default:
		color.PrintRedln("Unknown command: " + args[0])

	case "":
		// Do nothing.

	case "help", "h":
		cmds.PrintHelp()

	case "r":
		cmds.Read(buf, args[1:])

	case "clear", "clr":
		cmds.Clear()
	
	}
}
