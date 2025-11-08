package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
	"github.com/Moritisimor/Neo-Ed/internal/sigwatchers"
	"github.com/Moritisimor/Neo-Ed/internal/dispatch"
)

func main() {
	if len(os.Args) < 2 {
		color.PrintBlueln("Usage: ned <Target File>")
		color.PrintGreenln("Or run the 'h' command!")
		return
	}

	fileName := os.Args[1]
	openedFile := helpers.EnsureExistance(fileName)
	if helpers.IsDir(openedFile) {
		color.PrintRedln(fmt.Sprintf("Cannot Open '%s'! (Is a Directory)", fileName))
		return
	}

	writeBuf, initSize := helpers.ReadFileToBuffer(openedFile)
	color.PrintBlueln(fmt.Sprintf("Opened %d lines.", initSize))

	sigwatchers.StartSigTermWatcher()
	reader := helpers.CreateReader(color.SprintBlue(fmt.Sprintf("[%s] Ned >> ", fileName)))

	for {
		rawCmd, readErr := reader.Readline()
		if readErr != nil {
			color.PrintBlueln("Input interrupted. Try again.")
			continue
		}

		parts := strings.Split(strings.TrimSpace(rawCmd), " ")
		if parts[0] == "q" {
			color.PrintBlueln("Bye!")
			return
		}

		dispatch.Pipe(&writeBuf, parts, fileName)
	}
}
