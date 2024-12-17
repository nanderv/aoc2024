package chal5

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Bfunc(file io.Reader) any {
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

		if !hit {
			for {
				hit, pos1, pos2 := manualFulfillsRules(manual, manualAsMap, afterMap)
				if hit {
					break
				}
				// naively swap the positions of the elements of the failed rule
				a := manual[pos1]
				b := manual[pos2]
				manualAsMap[a] = pos2
				manualAsMap[b] = pos1
				manual[pos1] = b
				manual[pos2] = a
			}
			res += manual[len(manual)/2]
		}
	}
	return res
}
