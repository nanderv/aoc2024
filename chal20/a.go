package chal20

import (
	"bufio"
	"fmt"
	"github.com/nanderv/aoc2024/common"
	"io"
)

type Locater interface {
	Has(score inputScore) bool
}
type inputScore struct {
	common.Pos
	dist int
}

func Hash(score inputScore) int {
	return score.X*10000 + score.Y
}
func (score *inputScore) NB(mp [][]byte) []inputScore {
	ds := []common.Pos{{X: 1}, {X: -1}, {Y: 1}, {Y: -1}}
	res := []inputScore{}
	for _, d := range ds {
		p := score.Pos.Add(d)
		if mp[p.Y][p.X] != '#' {
			res = append(res, inputScore{Pos: p, dist: score.dist + 1})
		}
	}
	return res
}
func Afunc(file io.Reader) any {
	return AfuncT(file, 100)
}
func AfuncT(file io.Reader, minCheatWin int) any {
	scanner := bufio.NewScanner(file)
	mp := make([][]byte, 0)
	p := common.Pos{}
	e := common.Pos{}
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == 'S' {
				p.X = i
				p.Y = len(mp)
			}
			if c == 'E' {
				e.X = i
				e.Y = len(mp)
			}
		}
		if len(line) == 0 {
			break
		}
		mp = append(mp, []byte(line))
	}
	routeMap := common.NewMiMap(Hash)
	iQueue := []inputScore{{p, 0}}
	for len(iQueue) > 0 {
		nd := iQueue[0]
		iQueue = iQueue[1:]
		nxs := nd.NB(mp)

		for _, n := range nxs {
			if !routeMap.Has(n) {
				iQueue = append(iQueue, n)
				routeMap.Set(n)
			}
		}

	}
	res := 0
	for _, cheatStart := range routeMap.Keys() {
		dds := []common.Pos{{X: 2}, {X: 1, Y: 1}, {X: 1, Y: -1}, {Y: 2}, {Y: -2}, {X: -1, Y: 1}, {X: -1, Y: -1}, {X: -2}}
		for _, d := range dds {
			shortCut := inputScore{Pos: cheatStart.Pos.Add(d), dist: cheatStart.dist + 2}
			if routeMap.Has(shortCut) {
				target := routeMap.Get(shortCut)
				if target.dist-shortCut.dist >= minCheatWin {
					fmt.Println(cheatStart, shortCut, target.dist-shortCut.dist, res)

					res += 1
				}
			}
		}
	}
	return res
}
