package cmds

import (
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func Replace(buf *[]string, args []string, modified *bool) {
	if len(args) < 2 {
		color.PrintRedln("Usage: p <replacer> <replacee>")
		return
	}

	*modified = true

	replacee := args[0]
	replacer := args[1]

	for i, l := range *buf {
		temp := []string{}
		for i := range strings.SplitSeq(l, " ") {
			if i == replacee {
				temp = append(temp, replacer)
			} else {
				temp = append(temp, i)
			}
		}

		(*buf)[i] = strings.Join(temp, " ")
	}
}
