package cmds

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
)

func Insert(buf *[]string, args []string) {
	if len(args) < 1 {
		color.PrintRedln("Usage: i <Line>")
		return
	}

	index, err := strconv.ParseInt(args[0], 0, 32)
	if err != nil {
		color.PrintRedln(fmt.Sprintf("Expected a number, got '%s' instead.", args[0]))
		return
	}

	if int(index) > len(*buf) {
		color.PrintRedln("Invalid Index, this line does not exist in this file.")
		return
	}

	r := helpers.CreateReader(color.SprintMagenta(fmt.Sprintf("INSERT %d >> ", index)))
	lines := helpers.StartWriteLoop(r)

	*buf = slices.Insert(*buf, int(index) - 1, lines...)
}
