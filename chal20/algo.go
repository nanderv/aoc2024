package chal20

import (
	"bufio"
	"github.com/nanderv/aoc2024/common"
	"io"
)

func Algo(file io.Reader, minCheatWin int, mcc int) any {
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
	sn := inputScore{p, 0}
	iQueue := []inputScore{sn}
	routeMap.Set(sn)

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

	nbs := make([][]common.Pos, 0)
	for i := 2; i <= mcc; i++ {
		nbs = append(nbs, GetNbs(i))
	}

	res := 0
	r := common.NewMiMap(CHash)

	for _, cheatStart := range routeMap.Keys() {
		for i, dds := range nbs {
			for _, d := range dds {
				shortCut := inputScore{Pos: cheatStart.Pos.Add(d), dist: cheatStart.dist + i + 2}
				if routeMap.Has(shortCut) {
					target := routeMap.Get(shortCut)
					sco := target.dist - shortCut.dist
					if sco >= minCheatWin && !r.Has(rt{a: cheatStart, b: shortCut}) {

						r.Set(rt{a: cheatStart, b: shortCut})
						res += 1
					}
				}
			}
		}
	}

	return res
}
func GetNbs(i int) []common.Pos {
	res := make([]common.Pos, 0)
	for j := 0; j < i; j++ {
		res = append(res, common.Pos{X: j, Y: i - j})
		res = append(res, common.Pos{X: i - j, Y: -j})
		res = append(res, common.Pos{X: -j, Y: j - i})
		res = append(res, common.Pos{X: j - i, Y: j})
	}
	return res
}

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

type rt struct {
	a inputScore
	b inputScore
}

func CHash(r rt) int {
	a := r.a
	b := r.b
	z := 10000
	return a.X*z*z*z + a.Y*z*z + b.X*z + b.Y
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
