package chal6

import "fmt"

type pos struct {
	x int
	y int
	d int
}

func (p pos) fwd() pos {
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

func (p pos) rev() pos {
	if p.d == 0 {
		p.y += 1
	}
	if p.d == 1 {
		p.x -= 1
	}
	if p.d == 2 {
		p.y -= 1
	}
	if p.d == 3 {
		p.x += 1
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
func (p pos) GetIntegerHash(sy int) int {
	return 4*p.y + 4*sy*p.x + p.d
}

func next(d pos, base [][]byte) (pos, bool) {
	// look ahead
	dd := d.fwd()
	if dd.y < 0 || dd.x < 0 || dd.y >= len(base) || dd.x >= len(base[dd.y]) {
		return d, true
	}
	if base[dd.y][dd.x] == '#' {
		return d.turn(), false
	}

	return dd, false
}

func findPath(d pos, base [][]byte) map[string]pos {
	found := make(map[string]pos)
	found[d.String()] = d
	for {
		dd, done := next(d, base)
		if done {
			break
		}
		if _, ok := found[dd.String()]; !ok {
			found[dd.String()] = dd
			base[dd.y][dd.x] = 'X'
		}

		d = dd
	}
	return found
}
