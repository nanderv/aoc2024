package chal11

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type inp struct {
	num   int64
	count int32
}

var cache map[inp]int64

func init() {
	cache = make(map[inp]int64)
}
func Must[T any](a T, err error) T {
	if err != nil {
		panic(err)
	}
	return a
}
func Afunc(file io.Reader) int {

	scanner := bufio.NewScanner(file)
	numStr := make([]string, 0)

	for scanner.Scan() {
		numStr = append(numStr, strings.Fields(scanner.Text())...)
	}
	res := int64(0)
	fmt.Println(len(numStr))
	ops.Store(25)
	for _, num := range numStr {
		res += runAndCount(int64(Must(strconv.ParseInt(num, 10, 32))), 25)
	}
	fmt.Println(res)
	return int(res)
}

func runAndCount(num int64, count int32) int64 {
	if count > 5 && num < 1000 {
		ii := inp{num, count}
		if v, ok := cache[ii]; ok {
			return v
		}
	}
	nm := ops.Load()

	if nm > count {
		fmt.Println(count)
		ops.Store(count)
	}
	if count == 0 {
		return 1
	}
	v := int64(0)
	if num == 0 {
		v = runAndCount(1, count-1)
	} else if a, b, ok := numSplittableAndSplit(num); ok {
		v = runAndCount(a, count-1) + runAndCount(b, count-1)
	} else {
		v = runAndCount(num*2024, count-1)
	}
	if count > 5 && num < 1000 {
		ii := inp{num, count}
		cache[ii] = v
	}
	return v
}

func numSplittableAndSplit(num int64) (int64, int64, bool) {
	dg := num
	digitCount := 1
	for dg > 9 {
		dg = dg / 10
		digitCount++
	}
	if digitCount%2 == 1 {
		return 0, 0, false
	}

	pw := int64(math.Pow10(int(digitCount / 2)))

	return num / pw, num % pw, true
}
