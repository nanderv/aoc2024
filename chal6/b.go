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
		dd := zz.rev()
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

	for i := 0; i < (len(base)-2)*(len(base[0])-2); i++ {
		var done bool
		d, done = next(d, base)
		if done {
			return false
		}
	}

	return true
}
