package chal10

import (
	"bufio"
	"io"
)

func Bfunc(file io.Reader) any {
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

	res := 0
	for _, st := range starts {
		todo := []pos{st}
		for {
			if len(todo) == 0 {
				break
			}
			nxt := todo[0]
			if lines[nxt.y][nxt.x] == '9' {
				res++
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
