package dispatch

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/cmds"
)

func Pipe(buf *[]string, args []string, fileName string, modified *bool) {
	switch args[0] {
	default:
		color.PrintRedln("Unknown command: " + args[0])

	case "":
		// Do nothing.

	case "help", "h":
		cmds.PrintHelp()

	case "r":
		cmds.Read(buf, args[1:])

	case "a":
		cmds.Append(buf, modified)

	case "w":
		cmds.Write(buf, fileName, modified)

	case "i":
		cmds.Insert(buf, args[1:], modified)

	case "e":
		cmds.Edit(buf, args[1:], modified)

	case "f":
		cmds.Find(buf, args[1:])

	case "d":
		cmds.Delete(buf, args[1:], modified)

	case "x":
		cmds.Execute(args[1:])

	case "clear", "clr":
		cmds.Clear()
	}
}
