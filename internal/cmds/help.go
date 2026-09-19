package cmds

import "github.com/Moritisimor/EpsilonFetch/pkg/color"

func PrintHelp() {
	color.PrintRainbowln("==== HELP ====")

	color.PrintBlueln("Neo-Ed is an unofficial evolution of the original UNIX Text Editor ed.")
	color.PrintBlueln("It is implemented from scratch and written in Go.")
	color.PrintBlueln("Hint: <Parameter> means required, ?<Parameter>? means optional.")
	color.PrintBlueln("For more info visit https://github.com/Moritisimor/Neo-Ed")

	color.PrintBlueln("\n== a ==")
	color.PrintGreenln("This command will append Text to the tail of a file.")
	color.PrintMagentaln("Usage: a")

	color.PrintBlueln("\n== r ==")
	color.PrintGreenln("This command will read Text of a file. You can use it without parameters to read the entire file.")
	color.PrintGreenln("You can also read only a specific line by passing the line you want to read as an argument.")
	color.PrintGreenln("It is also possible to read ranges of lines by passing two arguments, the first being the start and the second being the end.")
	color.PrintMagentaln("Usage: r ?<Start>? ?<End>?")

	color.PrintBlueln("\n== w ==")
	color.PrintGreenln("This command is used for writing the internal buffer of the editor to the opened file.")
	color.PrintMagentaln("Usage: w")

	color.PrintBlueln("\n== x ==")
	color.PrintGreenln("This command can be used for executing programs from within the editor.")
	color.PrintMagentaln("Usage: x <Program> ?<Arguments>?")

	color.PrintBlueln("\n== f ==")
	color.PrintGreenln("This command is used for finding text within a file.")
	color.PrintMagentaln("Usage: f <Text>")

	color.PrintBlueln("\n== p ==")
	color.PrintGreenln("This command is used for replacing text with other text.")
	color.PrintMagentaln("Usage: p <TextToReplace> <TextToReplaceWith>")

	color.PrintBlueln("\n== i ==")
	color.PrintGreenln("This command can be used for inserting text between existing lines in a file.")
	color.PrintMagentaln("Usage: i <Line>")

	color.PrintBlueln("\n== e ==")
	color.PrintGreenln("This command is used for editing existing lines within a file.")
	color.PrintGreenln("It will only take one argument, which is the number of the line you want to edit.")
	color.PrintMagentaln("Usage: e <Line>")

	color.PrintBlueln("\n== d ==")
	color.PrintGreenln("This command is used for deleting existing lines.")
	color.PrintGreenln("It will take either one or two arguments, Those being either the line you want to delete or the range you want to delete.")
	color.PrintMagentaln("Usage: d <Start> ?<End>?")

	color.PrintBlueln("\n== q ==")
	color.PrintGreenln("This command is used for exiting Neo-Ed.")
	color.PrintMagentaln("Usage: q")

	color.PrintBlueln("\n== clear ==")
	color.PrintGreenln("This command is for clearing the screen of the terminal.")
	color.PrintMagentaln("Usage: clear")

	color.PrintRainbowln("==== HELP ====")
}
