package chal12

import (
	"bufio"
	"io"
)

type inp struct {
	num   int64
	count int32
}

var cache map[inp]int64

func init() {
	cache = make(map[inp]int64)
}
func Must[T any](a T, err error) T {
	if err != nil {
		panic(err)
	}
	return a
}

type pos struct {
	x int
	y int
}

func (p pos) add(q pos) pos {
	return pos{p.x + q.x, p.y + q.y}
}

func (p pos) allNeighbours(dirs []pos) []pos {
	pp := []pos{}
	for _, d := range dirs {
		r := p.add(d)
		pp = append(pp, r)
	}
	return pp
}

func getGroup(mp []string, dirs []pos, start pos) map[pos]struct{} {
	let := mp[start.y][start.x]
	found := make(map[pos]struct{})
	found[start] = struct{}{}

	allNgh := make([]pos, 1)
	allNgh[0] = start
	i := 0
	for i < len(allNgh) {
		p := allNgh[i]
		nb := p.allNeighbours(dirs)
		for _, n := range nb {
			if getV(mp, n) == let {
				if _, ok := found[n]; !ok {
					found[n] = struct{}{}
					allNgh = append(allNgh, n)
				}
			}
		}
		i++
	}

	return found
}

func Afunc(file io.Reader) int {
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
					nbs := p.allNeighbours(d)
					for _, nb := range nbs {
						if getV(mp, nb) != getV(mp, p) {
							wls++
						}
					}
					if len(nbs) < 4 {
						wls += 4 - len(nbs)
					}
				}
				res += wls * area
			}
		}
	}
	return res
}
