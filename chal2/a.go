package chal2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Afunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)

	res := 0
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		arr := make([]int, 0)

		for _, f := range strs {
			i, err := strconv.Atoi(f)
			if err != nil {
			} else {
				arr = append(arr, i)
			}
		}
		if len(arr) < 2 {
			res += 1
			continue
		}

		increasing := arr[0] < arr[1]
		safe := true
		for i := 1; i < len(arr); i++ {
			a := arr[i-1]
			b := arr[i]
			if !increasing {
				b, a = a, b
			}

			if b-a > 3 || b-a < 1 {
				safe = false
				break
			}
		}
		if safe {
			res += 1
		}
	}
	return res
}
