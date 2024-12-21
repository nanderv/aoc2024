package chal21

import (
	"bufio"
	"fmt"
	"io"
)

func Afunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		numPad, start := getMachine(2)
		n := getN(line)
		in := In(line, numPad, start)
		res += in * n
		fmt.Println(line, n, in, in*n)
	}
	return res
}

func getN(line string) int {
	mline := line
	for mline[0] == '0' {
		mline = mline[1:]
	}
	var ii int
	_, err := fmt.Sscanf(mline, "%dA", &ii)
	if err != nil {
		panic(err)
	}
	return ii
}
