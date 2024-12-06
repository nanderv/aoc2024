package chal6

import (
	"bufio"
	"io"
)

func Bfunc(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	base := make([][]byte, 0)
	i := 0
	d := pos{}
	for scanner.Scan() {
		ln := scanner.Text()
		lnn := []byte(ln)

		base = append(base, lnn)
		for p, l := range lnn {
			if l == '^' {
				d.x = p
				d.y = i
			}
		}
		i++
	}
	res := 0
	path := findPath(d, base)

	for _, zz := range path {

		pv := base[zz.y][zz.x]
		base[zz.y][zz.x] = '#'

		if calcLoop(d, base) {
			res += 1
		}
		base[zz.y][zz.x] = pv

	}
	return res
}

func calcLoop(d pos, base [][]byte) bool {
	res := 0
	found := make(map[string]any)
	res++
	for {
		// look ahead
		dd := d.nextSquare()
		if dd.y < 0 || dd.x < 0 || dd.y >= len(base) || dd.x >= len(base[dd.y]) {
			break
		}

		if base[dd.y][dd.x] == '#' {
			dd = d.turn()
		} else {
			_, ok := found[dd.String2()]
			if ok {
				return true
			}
			found[dd.String2()] = struct{}{}
		}
		d = dd
	}

	return false
}
