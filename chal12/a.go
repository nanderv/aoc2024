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

func (p pos) neighbours(dirs []pos, maxx, maxy int) []pos {
	pp := []pos{}
	for _, d := range dirs {
		r := p.add(d)
		if r.x >= 0 && r.y >= 0 && r.x < maxx && r.y < maxy {
			pp = append(pp, r)
		}
	}
	return pp
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
		nb := p.neighbours(dirs, len(mp[0]), len(mp))
		for _, n := range nb {
			if mp[n.y][n.x] == let {
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
