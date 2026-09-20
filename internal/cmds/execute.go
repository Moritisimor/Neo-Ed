package cmds

import (
	"fmt"
	"os"
	"os/exec"
)

func Execute(_ *EditorState, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: x <Command> ?<Args>?")
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
