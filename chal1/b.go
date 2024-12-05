package chal1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Bfunc(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	leftInput := make([]int, 0)
	rightInput := make([]int, 0)
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		i, err := strconv.Atoi(strs[0])
		if err != nil {
			fmt.Println(err)
		} else {
			leftInput = append(leftInput, i)
		}
		i, err = strconv.Atoi(strs[1])

		if err != nil {
			fmt.Println(err)
		} else {
			rightInput = append(rightInput, i)
		}
	}
	freqMap := make(map[int]int)
	for _, v := range rightInput {
		ii, ok := freqMap[v]
		if !ok {
			ii = 0
		}
		ii++
		freqMap[v] = ii
	}

	res := 0
	for _, v := range leftInput {
		ii, ok := freqMap[v]
		if !ok {
			ii = 0
		}
		res += v * ii
	}
	return res
}
