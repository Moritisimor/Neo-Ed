package cmds

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/Neo-Ed/internal/helpers"
)

func Append(buf *[]string) {
	r := helpers.CreateReader(color.SprintMagenta("APPEND >> "))
	lines := helpers.StartWriteLoop(r)
	(*buf) = append((*buf), lines...)
}
