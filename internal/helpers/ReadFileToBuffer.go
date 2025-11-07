package helpers

import (
	"bufio"
	"os"
)

func ReadFileToBuffer(file *os.File) (buffer []string, length int) {
	scanner := bufio.NewScanner(file)
	var buf []string
	l := 0

	for scanner.Scan() {
		l++
		buf = append(buf, scanner.Text())
	}

	return buf, l
}
