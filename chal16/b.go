package chal16

import (
	"bufio"
	"github.com/nanderv/aoc2024/common"
	"io"
	"sort"
)

func Bfunc(file io.Reader) int {
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

	found := NewMiMap(robot.Hash)
	nexts := []robot{{p: p}}
	for len(nexts) > 0 {
		nxt := nexts[0]

		nexts = nexts[1:]

		nx, ok, done := nxt.Fwd(mp)
		if done {
			return nx.score
		}
		if ok {
			if !found.Has(nx) {
				nexts = append(nexts, nx)
				found.Set(nx)
			} else {
				if found.Get(nx).score > nx.score {
					nexts = append(nexts, nx)
					found.Set(nx)
				}
			}
		}

		nx = nxt.Right(mp)
		if !found.Has(nx) {
			nexts = append(nexts, nx)
			found.Set(nx)
		} else {
			if found.Get(nx).score > nx.score {
				nexts = append(nexts, nx)
				found.Set(nx)
			}
		}

		nx = nxt.Left(mp)
		if !found.Has(nx) {
			nexts = append(nexts, nx)
			found.Set(nx)
		} else {
			if found.Get(nx).score > nx.score {
				nexts = append(nexts, nx)
				found.Set(nx)
			}
		}
		sort.Slice(nexts, func(i, j int) bool {
			return nexts[i].score < nexts[j].score
		})
	}
	panic("impossible")
}
