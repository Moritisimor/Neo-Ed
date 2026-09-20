package cmds

import (
	"fmt"
)

func Clear(_ *EditorState, _ []string) error {
	fmt.Print("\033[H\033[2J\033[3J")
	return nil
}
