package chal11

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync/atomic"
)

var ops atomic.Int32

func Bfunc(file io.Reader) int {

	scanner := bufio.NewScanner(file)
	numStr := make([]string, 0)

	for scanner.Scan() {
		numStr = append(numStr, strings.Fields(scanner.Text())...)
	}
	res := int64(0)
	ops.Store(76)
	fmt.Println("START", len(numStr))
	for _, num := range numStr {
		res += runAndCount(int64(Must(strconv.ParseInt(num, 10, 32))), 75)
	}
	fmt.Println(res)
	return int(res)
}
