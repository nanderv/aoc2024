package chal19

import (
	"bufio"
	"io"
	"strings"
)

type flag string

func Afunc(file io.Reader) any {
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
		if check(flag(word), flags) {
			res++
		}
	}
	return res
}
func check(f flag, m map[byte][]flag) bool {
	if len(f) == 0 {
		return true
	}

	if l, ok := m[f[0]]; ok {
		for _, fl := range l {
			if len(fl) > len(f) {
				continue
			}
			front := f[:len(fl)]
			if front == fl {
				if check(f[len(fl):], m) {
					return true
				}
			}
		}
	}
	return false
}
