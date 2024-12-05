package chal5

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Afunc(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	// map of int to list of ints that have to be after it
	afterMap := make(map[int][]int)
	for scanner.Scan() {
		var a, b int
		_, err := fmt.Sscanf(scanner.Text(), "%d|%d", &a, &b)
		if err != nil {
			break
		}
		lst, ok := afterMap[a]
		if !ok {
			lst = make([]int, 0)
		}
		lst = append(lst, b)
		afterMap[a] = lst
	}
	res := 0

	for scanner.Scan() {
		ln := scanner.Text()
		numsStr := strings.Split(ln, ",")
		// position map
		manual := make([]int, 0)
		manualAsMap := make(map[int]int, 0)
		for k, v := range numsStr {
			parseInt, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				panic(err)
			}
			manual = append(manual, int(parseInt))
			manualAsMap[int(parseInt)] = k
		}
		hit, _, _ := manualFulfillsRules(manual, manualAsMap, afterMap)

		if hit {
			res += manual[len(manual)/2]
		}
	}
	return res
}
