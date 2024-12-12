package chal12

import (
	"bufio"
	"io"
)

func Bfunc(file io.Reader) int {
	d := []pos{
		{x: 1, y: 0},
		{x: -1, y: 0},
		{x: 0, y: 1},
		{x: 0, y: -1},
	}
	scanner := bufio.NewScanner(file)
	mp := make([]string, 0)

	for scanner.Scan() {
		mp = append(mp, scanner.Text())
	}
	res := 0
	found := make(map[pos]struct{})
	for y, l := range mp {
		for x, _ := range l {
			if _, ok := found[pos{x, y}]; !ok {
				// create group
				gp := getGroup(mp, d, pos{x, y})
				area := len(gp)
				wls := 0
				for v, _ := range gp {
					found[v] = struct{}{}
					p := pos{v.x, v.y}
					nb := p.neighbours(d, len(mp[0]), len(mp))
					for _, n := range nb {
						if mp[n.y][n.x] != mp[y][x] {
							wls++
						}
					}
					if len(nb) < 4 {
						wls += 4 - len(nb)
					}
				}
				res += wls * area
			}
		}
	}
	return int(res)
}
