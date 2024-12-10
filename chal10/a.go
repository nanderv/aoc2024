package chal10

import (
	"bufio"
	"fmt"
	"io"
)

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
func Afunc(file io.Reader) int {
	d := []pos{
		{x: 1, y: 0},
		{x: -1, y: 0},
		{x: 0, y: 1},
		{x: 0, y: -1},
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	starts := make([]pos, 0)
	for y, l := range lines {
		for x, c := range l {
			if c == '0' {
				starts = append(starts, pos{x, y})
			}
		}
	}
	fmt.Println(starts)
	fmt.Println(d)
	res := 0
	for _, st := range starts {
		fmt.Println(st)
		todo := []pos{st}
		founds := make(map[pos]struct{})

		for {
			if len(todo) == 0 {
				break
			}
			nxt := todo[0]
			if lines[nxt.y][nxt.x] == '9' {
				fmt.Println(nxt)
				if _, ok := founds[nxt]; !ok {
					founds[nxt] = struct{}{}
					res++
				}
			}
			for _, nb := range nxt.neighbours(d, len(lines[0]), len(lines)) {
				if lines[nb.y][nb.x]-lines[nxt.y][nxt.x] == 1 {
					todo = append(todo, nb)
				}
			}

			todo = todo[1:]
		}

	}
	return res
}
