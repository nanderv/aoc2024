package chal21

import (
	"bufio"
	"github.com/nanderv/aoc2024/common"
	"io"
)

func Bfunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		numPad, start := getMachine(25)
		n := getN(line)
		in := In(line, numPad, start)
		res += in * n
	}
	return res
}

func getMachine(n int) (KeyPad, []common.Pos) {
	numPad, s1 := GetBaseKeypad()

	start := []common.Pos{s1}
	prev := &numPad

	for i := 0; i < n; i++ {
		rr, ss := GetRemoteKeypad()
		start = append(start, ss)
		prev.PrevDo = common.Memoize(PressIn.Hash, rr.Press)
		rr.PrevDo = Last
		prev = &rr
	}

	start = append(start, common.Pos{})
	return numPad, start
}
