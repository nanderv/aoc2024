package chal7

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Afunc(file io.Reader) int {
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

		if recurseOp(inputs[0], inputs[1:], result) {
			res += result
		}
	}

	return res
}

func recurseOp(intermediate int, nextNums []int, expectedResult int) bool {
	nxt := nextNums[0]
	im1 := intermediate + nxt
	res2 := intermediate * nxt
	if len(nextNums) > 1 {
		if recurseOp(im1, nextNums[1:], expectedResult) {
			return true
		}
		if recurseOp(res2, nextNums[1:], expectedResult) {
			return true
		}
	}
	return im1 == expectedResult || res2 == expectedResult
}
