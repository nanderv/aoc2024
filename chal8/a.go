package chal8

import (
	"bufio"
	"io"
)

type pos struct {
	x int
	y int
}

func (p pos) boinkOver(p2 pos) (res pos) {
	dx := p.x - p2.x
	dy := p.y - p2.y
	res.x = p.x + dx
	res.y = p.y + dy
	return res
}

const SKIP_CHAR = '.'

func Afunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	res := 0
	xs := 0
	ys := 0
	mp := make(map[byte][]pos)
	found := make(map[pos]any)
	for scanner.Scan() {
		txt := []byte(scanner.Text())
		xs = len(txt)
		for x, lt := range txt {
			if lt != SKIP_CHAR {
				v, ok := mp[lt]
				if !ok {
					v = make([]pos, 0)
				}
				v = append(v, pos{x: x, y: ys})
				mp[lt] = v
			}
		}
		ys++
	}

	for _, yy := range mp {
		for _, p1 := range yy {
			for _, p2 := range yy {
				if p1 != p2 {
					pz := p1.boinkOver(p2)
					if pz.x < 0 || pz.y < 0 || pz.x >= xs || pz.y >= ys {
						continue
					}
					if _, ok := found[pz]; !ok {
						res++
						found[pz] = struct{}{}
					}
				}
			}
		}
	}

	return res
}
