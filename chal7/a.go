package chal7

import (
	"bufio"
	"io"
)

func Afunc(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	return len(scanner.Bytes())
}
