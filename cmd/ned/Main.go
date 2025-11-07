package main

import (
	"log"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/tokenizing"
	"github.com/chzyer/readline"
)

func main() {
	reader, err := readline.NewEx(&readline.Config {
		Prompt: color.SprintBlue("CMD >> "),
		InterruptPrompt: "^C",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		rawCmd, readErr := reader.Readline()
		if readErr != nil {
			log.Fatal(readErr.Error())
		}

		parts := strings.Split(rawCmd, " ")
		tokenizing.Dispatch(parts)
	}
}
