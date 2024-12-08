package chal7

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Bfunc(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	res := 0

	for scanner.Scan() {
		txt := scanner.Text()

		ls := strings.Split(txt, ": ")
		result, _ := strconv.Atoi(ls[0])
		inputs := make([]int, 0)

		strngs := strings.Split(ls[1], " ")
		for _, r := range strngs {
			rr, _ := strconv.Atoi(r)
			inputs = append(inputs, rr)
		}

		if RecurseOp2(inputs[0], inputs[1:], result) {
			res += result
		}
	}

	return res
}

func RecurseOp2(intermediate int, nextNums []int, expectedResult int) bool {
	nxt := nextNums[0]
	im1 := intermediate + nxt
	im2 := intermediate * nxt
	rr := intermediate * 10
	cter := 10
	for nxt >= cter {
		rr *= 10
		cter *= 10
	}
	im3 := rr + nxt
	if len(nextNums) > 1 {
		if im1 <= expectedResult && RecurseOp2(im1, nextNums[1:], expectedResult) {
			return true
		}
		if im2 <= expectedResult && RecurseOp2(im2, nextNums[1:], expectedResult) {
			return true
		}
		if im3 <= expectedResult && RecurseOp2(im3, nextNums[1:], expectedResult) {
			return true
		}
		return false
	}
	return im1 == expectedResult || im2 == expectedResult || im3 == expectedResult
}
