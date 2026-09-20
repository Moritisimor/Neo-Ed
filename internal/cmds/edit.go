package cmds

import (
	"fmt"
	"strconv"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
)

func Edit(state *EditorState, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: e <line>")
	}

	line, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		return fmt.Errorf("Expected a number, got '%s' instead", args[1])
	}

	if int(line) > len(state.Buffer) {
		return fmt.Errorf("Invalid Index, this line does not exist in this file.")

	}

	if line < 1 {
		return fmt.Errorf("Cannot access indices which are below 1!")
	}

	r := helpers.CreateReader(color.SprintMagenta(fmt.Sprintf("EDIT %d >> ", line)))
	r.WriteStdin([]byte(state.Buffer[line-1]))

	text, readErr := r.Readline()
	if readErr != nil {
		return fmt.Errorf("Input interrupted.")
	}

	state.Buffer[line-1] = text
	state.Modified = true
	return nil
}
