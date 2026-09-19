package helpers

import (
	"bufio"
	"os"
)

func ReadFileToBuffer(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var buf []string

	for scanner.Scan() {
		buf = append(buf, scanner.Text())
	}

	return buf
}
