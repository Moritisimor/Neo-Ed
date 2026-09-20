package cmds

import (
	"fmt"
	"strconv"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
)

func Read(state *EditorState, args []string) error {
	if len(state.Buffer) == 0 {
		color.PrintYellowln("Empty File!")
		return nil
	}

	if len(args) == 0 {
		for i, l := range state.Buffer {
			helpers.PrintFileLine(i + 1, l)
		}

		return nil
	}

	var rangeEnd, rangeStart int
	for i, arg := range state.Buffer {
		parsedArg, parseErr := strconv.ParseInt(arg, 0, 64)
		if parseErr != nil {
			return fmt.Errorf("Expected a number, got '%s' instead", arg)
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
		return fmt.Errorf("Index may not be smaller than 1!")
	}

	if rangeEnd < rangeStart {
		return fmt.Errorf("The range end may not be smaller than the range start!")
	}

	if len(state.Buffer) < rangeStart || len(state.Buffer) < rangeEnd {
		return fmt.Errorf("Invalid Index, this line does not exist in this file.")
	}

	for i := rangeStart - 1; i <= rangeEnd - 1; i++ {
		helpers.PrintFileLine(i + 1, state.Buffer[i])
	}

	return nil
}
