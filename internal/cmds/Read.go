package cmds

import (
	"fmt"
	"strconv"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func Read(buf*[]string, args []string) {
	if len(args) < 1 {
		for _, l := range(*buf) {
			fmt.Println(l)
		}
		return
	}

	var rangeEnd, rangeStart int
	for i, arg := range(args) {
		parsedArg, parseErr := strconv.ParseInt(arg, 0, 64)
		if parseErr != nil {
			color.PrintRedln("Expected integer, got '" + arg + "' instead.")
			return
		}

		if i == 0 {
			rangeStart = int(parsedArg)
			rangeEnd = int(parsedArg)
		}

		if i > 0 {
			rangeEnd = int(parsedArg)
		}
	}

	if rangeStart < 1 || rangeEnd < 1 {
		color.PrintRedln("Index may not be smaller than 1!")
		return
	}

	if rangeEnd < rangeStart {
		color.PrintRedln("The range end may not be smaller than the range start!")
		return
	}

	if len(*buf) < rangeStart || len(*buf) < rangeEnd {
		color.PrintRedln("Invalid Index, this line does not exist in this file.")
		return
	}

	for i := rangeStart - 1; i <= rangeEnd - 1; i++ {
		fmt.Println((*buf)[i])
	}
}
