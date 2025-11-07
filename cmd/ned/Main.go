package main

import (
	"log"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/sigwatchers"
	"github.com/Moritisimor/Neo-Ed/internal/tokenizing"
	"github.com/chzyer/readline"
)

func main() {
	sigwatchers.StartSigTermWatcher()

	

	reader, creationErr := readline.NewEx(&readline.Config {
		Prompt: color.SprintBlue("CMD >> "),
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

		tokenizing.Dispatch(parts)
	}
}
