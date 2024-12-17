package chal2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Bfunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)

	res := 0
	ii := 0
	for scanner.Scan() {
		ii++
		strs := strings.Fields(scanner.Text())
		arr := make([]int, 0)

		for _, f := range strs {
			i, err := strconv.Atoi(f)
			if err != nil {
			} else {
				arr = append(arr, i)
			}
		}

		for i := 0; i < len(arr); i++ {
			num := deleteAtIndex(arr, i)
			safe := calcForSlice(num)
			if safe {
				res += 1
				break
			}
		}

	}
	return res
}

func calcForSlice(arr []int) bool {
	increasing := arr[0] < arr[1]
	for i := 1; i < len(arr); i++ {
		a := arr[i-1]
		b := arr[i]
		if !increasing {
			b, a = a, b
		}
		if b-a > 3 || b-a < 1 {
			return false
		}
	}
	return true
}

func deleteAtIndex[T any](slice []T, i int) []T {
	return append(slice[:i:i], slice[i+1:]...)
}
