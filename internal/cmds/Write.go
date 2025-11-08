package cmds

import (
	"fmt"
	"os"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func Write(buf *[]string, fileName string) {
	bufAsString := strings.Join((*buf), "\n")
	err := os.WriteFile(fileName, []byte(bufAsString), 0755)
	if err != nil {
		color.PrintRedln(fmt.Sprintf("Writing buffer to '%s' failed!\nError: %s", fileName, err.Error()))
		return
	}

	color.PrintGreen(fmt.Sprintf("Successfully wrote buffer to '%s'", fileName))
}
