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
					pValue := getV(mp, p)
					nbs := p.allNeighbours(d)
					for _, nb := range nbs {
						nbV := getV(mp, nb)
						if nbV != pValue {
							if nb.y != p.y {
								// slide over x axis
								pSlide := p.add(pos{x: -1, y: 0})
								nbSlide := nb.add(pos{x: -1, y: 0})
								// first condition is : if one to the left is same value, it's same wall
								// second condition is: inner corner check
								if getV(mp, pSlide) != pValue || getV(mp, nbSlide) == pValue {
									wls++
								}
							}
							if nb.x != p.x {
								// slide over y axis
								pSlide := p.add(pos{x: 0, y: -1})
								nbSlide := nb.add(pos{x: 0, y: -1})
								// first condition is : if one upwards is same value, it's same wall
								// second condition is: inner corner check
								if getV(mp, pSlide) != pValue || getV(mp, nbSlide) == pValue {
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
	return res
}
