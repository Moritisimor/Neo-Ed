package cmds

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
)

func Append(state *EditorState, _ []string) error {
	state.Modified = true
	r := helpers.CreateReader(color.SprintMagenta("APPEND >> "))
	lines := helpers.StartWriteLoop(r)
	state.Buffer = append(state.Buffer, lines...)
	return nil
}
