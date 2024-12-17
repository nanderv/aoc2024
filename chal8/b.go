package chal8

import (
	"bufio"
	"io"
)

func Bfunc(file io.Reader) any {
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
					if _, ok := found[p1]; !ok {
						res++
						found[p1] = struct{}{}
					}
					if _, ok := found[p2]; !ok {
						res++
						found[p2] = struct{}{}
					}
					pz := p1
					py := p2
					for {
						prev := pz
						pz = pz.boinkOver(py)
						py = prev
						if pz.x < 0 || pz.y < 0 || pz.x >= xs || pz.y >= ys {
							break
						}
						if _, ok := found[pz]; !ok {
							res++
							found[pz] = struct{}{}
						}
					}
				}
			}
		}
	}

	return res
}
