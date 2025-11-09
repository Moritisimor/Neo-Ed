package cmds

import (
	"fmt"
	"strconv"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func Delete(buf *[]string, args []string) {
	if len(args) < 1 {
		color.PrintRedln("Usage: d <Line> || d <From> <To>")
	}

	index, err := strconv.ParseInt(args[0], 0, 32)
	if err != nil {
		color.PrintRedln(fmt.Sprintf("Expected a number, got '%s' instead.", args[0]))
	}

	if index < 1 {
		color.PrintRedln("Index may not be smaller than 1!")
		return
	}

	if len(*buf) < int(index) {
		color.PrintRedln("Invalid Index, this line does not exist in this file.")
		return
	}

	(*buf) = append((*buf)[:index - 1], (*buf)[index:]...)
}
