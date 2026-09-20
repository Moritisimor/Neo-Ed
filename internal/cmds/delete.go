package cmds

import (
	"fmt"
	"slices"
	"strconv"
)

func Delete(state *EditorState, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Invalid args! Usage: d <Line>")
	}

	var rangeEnd, rangeStart int
	for i, arg := range args {
		num, err := strconv.ParseInt(arg, 0, 32)
		if err != nil {
			return fmt.Errorf("Expected a number, got '%s' instead.", args[0])
		}

		if i == 0 {
			rangeEnd = int(num)
			rangeStart = int(num)
		} else {
			rangeEnd = int(num)
		}
	}

	if rangeEnd < 1 || rangeStart < 1 {
		return fmt.Errorf("Index may not be smaller than 1!")
	}

	if rangeEnd < rangeStart {
		return fmt.Errorf("The range end may not be smaller than the range start!")
	}

	if len(state.Buffer) < int(rangeStart) || len(state.Buffer) < int(rangeEnd) {
		return fmt.Errorf("Invalid Index, this line does not exist in this file.")
	}

	state.Buffer = slices.Delete(state.Buffer, rangeStart-1, rangeEnd)
	state.Modified = true
	return nil
}
