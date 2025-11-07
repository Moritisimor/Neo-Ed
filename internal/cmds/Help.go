package cmds

import "github.com/Moritisimor/EpsilonFetch/pkg/color"

func PrintHelp() {
	color.PrintRainbowln("==== HELP ====")

	color.PrintBlueln("Neo-Ed is an unofficial evolution of the original UNIX Text Editor ed.")
	color.PrintGreenln("It is implemented from scratch and written in Go.")
	color.PrintGreenln("So far, as this is the earliest version, no commands have been implemented yet.")

	color.PrintRainbowln("==== HELP ====")
}
