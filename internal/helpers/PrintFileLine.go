package helpers

import "fmt"

func PrintFileLine(line int, lineContent string) {
	if line < 10 {
		fmt.Printf("%d     | %s\n", line, lineContent)
		return
	}

	if line < 100 {
		fmt.Printf("%d    | %s\n", line, lineContent)
		return
	}
	
	if line < 1000 {
		fmt.Printf("%d   | %s\n", line, lineContent)
		return
	}

	if line < 10000 {
		fmt.Printf("%d  | %s\n", line, lineContent)
		return
	}

	if line < 100000 {
		fmt.Printf("%d | %s\n", line, lineContent)
		return
	}
}
