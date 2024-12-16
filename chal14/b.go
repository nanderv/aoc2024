package chal14

import (
	"bufio"
	"fmt"
	"github.com/nanderv/aoc2024/common"
	"io"
)

func Bfunc(file io.Reader) int {
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
	pps := make([]common.Pos, len(robots))
	for n := 1; true; n++ {
		for i, r := range robots {
			r.move(mapBounds.X, mapBounds.Y, 1)
			robots[i] = r
			pps[i] = r.p
		}
		tr := isTree(robots)
		if tr {
			common.VisualisePos(pps, fmt.Sprintf("%d.png", n))
			return n
		}
		if n > 10000 {
			return -100
		}
	}
	panic("Impossibru")
}

func isTree(robots []robot) bool {
	ps := make(map[common.Pos]struct{})
	for _, r := range robots {
		ps[r.p] = struct{}{}
	}

	for _, r := range robots {
		found := true
		for i := 0; i < 30; i++ {
			z := r.p.Add(common.Pos{Y: 1}.Mul(i))
			if _, ok := ps[z]; !ok {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}

	return false
}
