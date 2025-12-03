package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/dispatch"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
	"github.com/Moritisimor/Neo-Ed/internal/cmds"
)

// Neo-Ed is an ed-like text-editor, and this is its main package.
func main() {
	if len(os.Args) < 2 {
		color.PrintBlueln("Usage: ned <Target File>")
		color.PrintGreenln("Or run the 'h' command!")
		return
	}

	fileName := os.Args[1]
	openedFile := helpers.EnsureExistence(fileName)
	if helpers.IsDir(openedFile) {
		color.PrintRedln(fmt.Sprintf("Cannot Open '%s'! (Is a Directory)", fileName))
		return
	}


	writeBuf := helpers.ReadFileToBuffer(openedFile)
	color.PrintBlueln(fmt.Sprintf("Opened %d lines.", len(writeBuf)))
	modified := false
	reader := helpers.CreateReader(color.SprintBlue(fmt.Sprintf("[%s] Ned >> ", fileName)))

	for {
		rawCmd, readErr := reader.Readline()
		if readErr != nil {
			color.PrintBlueln("Input interrupted. Try again.")
			continue
		}

		parts := strings.Split(strings.TrimSpace(rawCmd), " ")
		if parts[0] == "q" {
			if modified {
				color.PrintRedln("Cannot exit as the buffer is modified. Save changes using 'w' or force quit using 'q!'")
				continue
			} else {
				color.PrintBlueln("Bye!")
				return
			}
		}

		if parts[0] == "q!" {
			color.PrintBlueln("Bye!")
			return
		}

		if parts[0] == "wq" {
			if cmds.Write(&writeBuf, fileName, &modified) != nil {
				continue
			}

			color.PrintBlueln("Bye!")
			return
		}

		dispatch.Pipe(&writeBuf, parts, fileName, &modified)
	}
}
