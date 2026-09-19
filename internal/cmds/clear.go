package cmds

import "fmt"

func Clear() {
	fmt.Print("\033[H\033[2J\033[3J")
}
