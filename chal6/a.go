package chal6

import (
	"bufio"
	"fmt"
	"io"
)

type pos struct {
	x int
	y int
	d int
}

func (p pos) nextSquare() pos {
	if p.d == 0 {
		p.y -= 1
	}
	if p.d == 1 {
		p.x += 1
	}
	if p.d == 2 {
		p.y += 1
	}
	if p.d == 3 {
		p.x -= 1
	}
	return p
}
func (p pos) turn() pos {
	p.d += 1
	if p.d > 3 {
		p.d = 0
	}
	return p
}
func (p pos) String() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}

func (p pos) String2() string {
	return fmt.Sprintf("%d:%d:%d", p.x, p.y, p.d)
}
func Afunc(file io.Reader) int {
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
	found := make(map[string]any)
	found[d.String()] = struct{}{}
	res++
	for {
		// look ahead
		dd := d.nextSquare()
		if dd.y < 0 || dd.x < 0 || dd.y >= len(base) || dd.x >= len(base[dd.y]) {
			break
		}

		if base[dd.y][dd.x] == '#' {
			dd = d.turn()
		} else {

			_, ok := found[dd.String()]
			if !ok {
				res++
			}
			found[dd.String()] = struct{}{}
			base[dd.y][dd.x] = 'X'
		}
		d = dd
	}
	for _, ln := range base {
		fmt.Println(string(ln))
	}

	return res
}
