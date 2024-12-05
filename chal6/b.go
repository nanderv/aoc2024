package chal6

import (
	"bufio"
	"io"
)

func Bfunc(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	// map of int to list of ints that have to be after it
	return len(scanner.Bytes())

}
