package chal22

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Afunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		sec, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		for i := 0; i < 2000; i++ {
			sec = Nxt(sec)
		}
		res += sec
	}
	return res
}

func Nxt(sec int) int {
	im := sec * 64
	sec = sec ^ im
	sec = sec % 16777216

	im = sec / 32
	sec = sec ^ im
	sec = sec % 16777216

	im = sec * 2048
	sec = sec ^ im
	sec = sec % 16777216
	return sec
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
