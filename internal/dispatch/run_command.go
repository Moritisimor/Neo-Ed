package dispatch

import (
	"fmt"

	"github.com/Moritisimor/Neo-Ed/internal/cmds"
)

func RunCommand(editorState *cmds.EditorState, args []string) error {
	fun := editorState.CommandRegistry[args[0]]
	if fun == nil {
		return fmt.Errorf("No such command: %s", args[0])
	}

	return fun(editorState, args[1:])
}
