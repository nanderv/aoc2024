package chal16

import (
	"bufio"
	"fmt"
	"github.com/nanderv/aoc2024/common"
	"io"
)

func Bfunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	mp := make([][]byte, 0)
	p := common.Pos{}
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == 'S' {
				p.X = i
				p.Y = len(mp)
			}
		}
		if len(line) == 0 {
			break
		}
		mp = append(mp, []byte(line))
	}

	i, done := getResult(p, mp)
	if !done {
		panic("impossible")
	}
	found := common.NewMiMap(robot.Hash)
	foundTraces := make([]robot, 0)
	nexts := []robot{{p: p}}

	for len(nexts) > 0 {

		next := nexts[0]
		nexts = nexts[1:]
		if found.Has(next) {
			if found.Get(next).score < next.score {
				continue
			}
		}
		found.Set(next)

		if next.score > i {
			continue
		}

		nx, ok, done := next.Fwd(mp)
		if done {
			if nx.score == i {
				foundTraces = append(foundTraces, next)
			}
			continue
		}
		if ok {
			nexts = append(nexts, nx)
		}

		nx2 := next.Right(mp)
		nexts = append(nexts, nx2)

		nx3 := next.Left(mp)
		nexts = append(nexts, nx3)

	}
	found = common.NewMiMap(robot.PosHash)
	fmt.Println("t", len(foundTraces))
	for _, rob := range foundTraces {
		rr := rob
		for true {
			found.Set(rr)
			rz := rr.prev
			if rz == nil {
				break
			}
			rr = *rz
		}
	}

	return found.Len() + 1
}
