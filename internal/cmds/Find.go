package cmds

import (
	"fmt"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
)

func Find(buf *[]string, args []string) {
	if len(args) < 1 {
		color.PrintRedln("Usage: f <text>")
		return
	}

	matches := 0
	text := strings.TrimSpace(strings.Join(args, " "))
	for i, line := range(*buf) {
		if strings.Contains(line, text) {
			color.PrintGreenln(fmt.Sprintf("Match in line %d", i + 1))
			helpers.PrintFileLine(i + 1, line)
			matches++
		}
	}

	if matches == 0 {
		color.PrintYellowln("No Matches!")
	} else {
		color.PrintGreenln(fmt.Sprintf("%d Matches.", matches))
	}
}
