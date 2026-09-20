package cmds

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
)

func Insert(state *EditorState, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: i <Line>")
	}

	index, err := strconv.ParseInt(args[0], 0, 32)
	if err != nil {
		return fmt.Errorf("Expected a number, got '%s' instead.", args[0])
	}

	if int(index) > len(state.Buffer) {
		return fmt.Errorf("Invalid Index, this line does not exist in this file.")
	}

	r := helpers.CreateReader(color.SprintMagenta(fmt.Sprintf("INSERT %d >> ", index)))
	lines := helpers.StartWriteLoop(r)

	state.Buffer = slices.Insert(state.Buffer, int(index) - 1, lines...)
	state.Modified = true
	return nil
}
