package cmds

import (
	"fmt"
	"strings"
)

func Replace(state *EditorState, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("Usage: p <replacer> <replacee>")
	}

	state.Modified = true

	replacee := args[0]
	replacer := args[1]

	for i, l := range state.Buffer {
		temp := []string{}
		for i := range strings.SplitSeq(l, " ") {
			if i == replacee {
				temp = append(temp, replacer)
			} else {
				temp = append(temp, i)
			}
		}

		state.Buffer[i] = strings.Join(temp, " ")
	}

	return nil
}
