package chal6

import (
	"bufio"
	"io"
)

func Bfunc(file io.Reader) int {
	//debug.SetGCPercent(-1)
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
		dd := zz.prevSquare()
		if calcLoop(dd, &base) {
			res += 1
		}
		base[zz.y][zz.x] = pv

	}
	return res
}

func calcLoop(d pos, bs *[][]byte) bool {
	base := *bs
	res := 0
	res++
	mp := make([]bool, len(base)*len(base[0])*4)
	for {
		// look ahead
		dd := d.nextSquare()
		if dd.y < 0 || dd.x < 0 || dd.y >= len(base) || dd.x >= len(base[dd.y]) {
			break
		}

		if base[dd.y][dd.x] == '#' {
			dd = d.turn()
		} else {
			if mp[dd.GetIntegerHash(len(base))] {
				return true
			}
			mp[dd.GetIntegerHash(len(base))] = true
		}
		d = dd
	}

	return false
}
