package chal6

import (
	"bufio"
	"io"
)

func Afunc(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return len(scanner.Bytes())
}
