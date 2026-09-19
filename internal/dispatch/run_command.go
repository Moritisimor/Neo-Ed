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

	case "r", "read":
		cmds.Read(buf, args[1:])

	case "a", "append":
		cmds.Append(buf, modified)

	case "w", "write":
		cmds.Write(buf, fileName, modified)

	case "i", "insert":
		cmds.Insert(buf, args[1:], modified)

	case "e", "edit":
		cmds.Edit(buf, args[1:], modified)

	case "f", "find":
		cmds.Find(buf, args[1:])

	case "p", "replace":
		cmds.Replace(buf, args[1:], modified)

	case "d", "delete":
		cmds.Delete(buf, args[1:], modified)

	case "x", "execute":
		cmds.Execute(args[1:])

	case "clear", "clr", "c":
		cmds.Clear()
	}
}
