package cmds

import (
	"fmt"
	"os"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func Write(state *EditorState, _ []string) error {
	bufAsString := strings.Join(state.Buffer, "\n")
	err := os.WriteFile(state.FileName, []byte(bufAsString), 0755)
	if err != nil {
		return fmt.Errorf("Writing buffer to '%s' failed!\nError: %s", state.FileName, err.Error())
	}

	color.PrintGreenln(fmt.Sprintf("Successfully wrote buffer to '%s'", state.FileName))
	state.Modified = false
	return nil
}
