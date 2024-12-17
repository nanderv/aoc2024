package chal1

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func Afunc(file io.Reader) any {
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
	slices.Sort(leftInput)
	slices.Sort(rightInput)

	res := 0
	for i := 0; i < len(leftInput); i++ {
		v := leftInput[i] - rightInput[i]
		if v < 0 {
			v = -v
		}
		res += v
	}
	return res
}
