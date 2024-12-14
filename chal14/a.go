package chal14

import (
	"bufio"
	"fmt"
	"github.com/nanderv/aoc2024/common"
	"io"
)

type robot struct {
	p common.Pos
	v common.Pos
}

func (r *robot) move(maxX int, maxY int, times int) {
	for i := 0; i < times; i++ {
		p := r.p.Add(r.v)

		if p.X < 0 {
			p.X += maxX
		}
		if p.Y < 0 {
			p.Y += maxY
		}
		if p.X >= maxX {
			p.X -= maxX
		}
		if p.Y >= maxY {
			p.Y -= maxY
		}
		r.p = p
	}
}
func Afunc(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	mapBounds := common.Pos{}
	fmt.Sscanf(scanner.Text(), "map X=%d Y=%d", &mapBounds.X, &mapBounds.Y)
	robots := make([]robot, 0)
	for scanner.Scan() {
		r := robot{}
		fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &r.p.X, &r.p.Y, &r.v.X, &r.v.Y)
		robots = append(robots, r)
	}
	// simulate
	for i, r := range robots {
		r.move(mapBounds.X, mapBounds.Y, 100)
		robots[i] = r
	}
	xB := (mapBounds.X - 1) / 2
	yB := (mapBounds.Y - 1) / 2
	fmt.Println(mapBounds.X, mapBounds.Y, xB, yB)
	r0 := 0
	r1 := 0
	r2 := 0
	r3 := 0
	r4 := 0
	for _, r := range robots {
		if r.p.X == xB || r.p.Y == yB {
			r0++
			continue
		}
		if r.p.X < xB && r.p.Y < yB {
			r1++
			continue
		}

		if r.p.X > xB && r.p.Y < yB {
			r2++
			continue
		}
		if r.p.X < xB && r.p.Y > yB {
			r3++
			continue
		}
		if r.p.X > xB && r.p.Y > yB {
			r4++
			continue
		}
		panic("Disco")
	}
	fmt.Println(len(robots), r1, r2, r3, r4, r1+r2+r3+r4+r0)
	return r1 * r2 * r3 * r4
}
