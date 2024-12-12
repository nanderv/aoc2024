package chal12

import (
	"bufio"
	"io"
)

func getV(mp []string, p pos) byte {
	if p.x < 0 || p.y < 0 || p.x >= len(mp[0]) || p.y >= len(mp) {
		return '$'
	}
	return mp[p.y][p.x]
}
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
				for p, _ := range gp {
					found[p] = struct{}{}
					pV := getV(mp, p)
					nb := p.allNeighbours(d)
					for _, n := range nb {
						vv := getV(mp, n)
						if vv != pV {
							if n.y != p.y {
								// move over x axis
								m := p.add(pos{x: -1, y: 0})
								mm := n.add(pos{x: -1, y: 0})
								if getV(mp, m) != pV || getV(mp, mm) == pV {
									wls++
								}
							}
							if n.x != p.x {
								// move over y axis

								m := p.add(pos{x: 0, y: -1})
								mm := n.add(pos{x: 0, y: -1})
								if getV(mp, m) != pV || getV(mp, mm) == pV {
									wls++
								}
							}
						}
					}
				}
				res += wls * area
			}
		}
	}
	return int(res)
}
