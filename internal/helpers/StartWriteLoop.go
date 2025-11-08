package helpers

import (
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/chzyer/readline"
)

func StartWriteLoop(r *readline.Instance) []string {
	var lines []string

	for {
		tempbuf, err := r.Readline()
		if err != nil {
			color.PrintRedln("Input interrupted. Try again.")
			return []string {}
		}

		if strings.TrimSpace(tempbuf) == "." {
			break
		}

		lines = append(lines, tempbuf)
	}

	return lines
}