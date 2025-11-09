package cmds

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func Execute(args []string) {
	if len(args) < 2 {
		color.PrintRedln("Usage: x <Command> ?<Args>?")
		return
	}

	cmd := exec.Command(args[1], args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	
	runErr := cmd.Run()
	if runErr != nil {
		color.PrintRedln(fmt.Sprintf("Executing '%s' failed.\nError: %s", args[1], runErr.Error()))
	}
}
