package chal19

import (
	"bufio"
	"io"
	"strings"
)

func Bfunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	str := strings.Replace(scanner.Text(), " ", "", -1)
	flStr := strings.Split(str, ",")

	flags := make(map[byte][]flag)
	for _, fl := range flStr {
		if _, ok := flags[fl[0]]; !ok {
			flags[fl[0]] = make([]flag, 0)
		}
		flags[fl[0]] = append(flags[fl[0]], flag(fl))
	}

	scanner.Scan()
	res := 0
	for scanner.Scan() {
		word := scanner.Text()
		res += checkB(flag(word), flags)
	}
	return res
}

func checkB(f flag, m map[byte][]flag) int {
	res := make([]int, len(f)+1)
	res[0] = 1
	for i := 0; i < len(f); i++ {
		if l, ok := m[f[i]]; ok {
			for _, fl := range l {
				if len(fl)+i > len(f) {
					continue
				}
				front := f[i : i+len(fl)]
				if front == fl {
					res[i+len(fl)] += res[i]
				}
			}
		}
	}
	return res[len(res)-1]
}
