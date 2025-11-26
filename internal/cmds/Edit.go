package cmds

import (
	"fmt"
	"strconv"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
)

func Edit(buf *[]string, args []string, modified *bool) {
	if len(args) < 1 {
		color.PrintBlueln("Usage: e <line>")
		return
	}

	line, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		color.PrintRedln(fmt.Sprintf("Expected a number, got '%s' instead", args[1]))
		return
	}

	if int(line) > len(*buf) {
		color.PrintRedln("Invalid Index, this line does not exist in this file.")
		return
	}

	if line < 1 {
		color.PrintRedln("Cannot access indices which are below 1!")
		return
	}

	r := helpers.CreateReader(color.SprintMagenta(fmt.Sprintf("EDIT %d >> ", line)))
	r.WriteStdin([]byte((*buf)[line-1]))

	text, readErr := r.Readline()
	if readErr != nil {
		color.PrintRedln("Input interrupted.")
		return
	}

	(*buf)[line-1] = text
	*modified = true
}
