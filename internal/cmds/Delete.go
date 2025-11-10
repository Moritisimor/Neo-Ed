package cmds

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func Delete(buf *[]string, args []string) {
	if len(args) < 1 {
		color.PrintRedln("Usage: d <Line> | d <From> <To>")
	}

	var rangeEnd, rangeStart int
	for i, arg := range(args) {
		num, err := strconv.ParseInt(arg, 0, 32)
		if err != nil {
			color.PrintRedln(fmt.Sprintf("Expected a number, got '%s' instead.", args[0]))
		}

		if i == 0 {
			rangeEnd = int(num)
			rangeStart = int(num)
		} else {
			rangeEnd = int(num)
		}
	}

	if rangeEnd < 1 || rangeStart < 1 {
		color.PrintRedln("Index may not be smaller than 1!")
		return
	}

	if rangeEnd < rangeStart {
		color.PrintRedln("The range end may not be smaller than the range start!")
		return
	}

	if len(*buf) < int(rangeStart) || len(*buf) < int(rangeEnd) {
		color.PrintRedln("Invalid Index, this line does not exist in this file.")
		return
	}

	*buf = slices.Delete(*buf, rangeStart - 1, rangeEnd)
}
