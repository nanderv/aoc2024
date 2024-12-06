package chal7

import (
	"bufio"
	"io"
)

func Bfunc(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	return len(scanner.Bytes())
}
