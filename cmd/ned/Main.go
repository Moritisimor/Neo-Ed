package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
	"github.com/Moritisimor/Neo-Ed/internal/sigwatchers"
	"github.com/Moritisimor/Neo-Ed/internal/dispatch"
	"github.com/chzyer/readline"
)

func main() {
	if len(os.Args) < 2 {
		color.PrintBlueln("Usage: ned <Target File>")
		color.PrintGreenln("Or run the 'h' command!")
		return
	}

	fileName := os.Args[1]
	openedFile, openErr := os.Open(fileName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	writeBuf, initSize := helpers.ReadFileToBuffer(openedFile)
	color.PrintBlueln(fmt.Sprintf("Opened %d lines.", initSize))

	sigwatchers.StartSigTermWatcher()
	reader, creationErr := readline.NewEx(&readline.Config {
		Prompt: color.SprintBlue(fmt.Sprintf("[%s] CMD >> ", fileName)),
		InterruptPrompt: "^C",
	})

	if creationErr != nil {
		log.Fatal(creationErr.Error())
	}

	for {
		rawCmd, readErr := reader.Readline()
		if readErr != nil {
			color.PrintBlueln("Input interrupted. Try again.")
			continue
		}

		parts := strings.Split(rawCmd, " ")
		if parts[0] == "q" {
			color.PrintBlueln("Bye!")
			return
		}

		dispatch.Pipe(&writeBuf, parts)
	}
}
