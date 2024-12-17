package chal15

import (
	"bufio"
	"github.com/nanderv/aoc2024/common"
	"io"
	"sort"
)

func BExec(mp [][]byte, p common.Pos, cmd string) int {
	for _, c := range cmd {
		mp[p.Y][p.X] = '.'
		if c == '^' {
			p = tryToMoveY(mp, p, common.Pos{Y: -1})
		}
		if c == 'v' {
			p = tryToMoveY(mp, p, common.Pos{Y: 1})
		}
		if c == '>' {
			p = tryToMoveY(mp, p, common.Pos{X: 1})
		}
		if c == '<' {
			p = tryToMoveY(mp, p, common.Pos{X: -1})
		}
		mp[p.Y][p.X] = '@'

	}
	res := 0
	for y, ln := range mp {
		for x, c := range ln {
			if c == 'O' || c == '[' {
				res += y*100 + x
			}
		}
	}
	return res

}
func Bfunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	mp := make([][]byte, 0)
	p := common.Pos{}
	needsExpanse := false
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == '@' {
				p.X = i
				p.Y = len(mp)
			}
			if c == 'O' {
				needsExpanse = true
			}
			if c == '[' && needsExpanse == true {
				panic("inconsistent")
			}
		}
		if len(line) == 0 {
			break
		}
		mp = append(mp, []byte(line))
	}
	if needsExpanse {
		mp = expand(mp)
		p = p.DirMul(common.Pos{X: 2, Y: 1})
	}
	str := ""

	for scanner.Scan() {
		ln := scanner.Text()
		str += ln

	}
	return BExec(mp, p, str)
}

func MergeMaps[T comparable, V any](a map[T]V, b map[T]V) map[T]V {
	res := make(map[T]V)
	for k, v := range a {
		res[k] = v
	}
	for k, v := range b {
		res[k] = v
	}
	return res
}

func getTheMoves(mp [][]byte, p common.Pos, d common.Pos, didNeighbour bool) (map[common.Pos]struct{}, bool) {
	or := p.Add(d)
	if mp[or.Y][or.X] == '.' {
		return map[common.Pos]struct{}{p: {}}, true
	}
	if mp[or.Y][or.X] == '#' {
		return nil, false
	}

	// Easy case: horizontal axis
	if mp[or.Y][or.X] == 'O' || ((d.Y == 0 || didNeighbour) && (mp[or.Y][or.X] == '[' || mp[or.Y][or.X] == ']')) {
		mv, cm := getTheMoves(mp, or, d, false)
		if !cm {
			return nil, false
		}
		mv[or] = struct{}{}
		return mv, true
	}

	// weird case: vertical axis,
	p2 := p
	act := false
	if mp[or.Y][or.X] == '[' {
		p2.X += 1
		act = true
	}
	if mp[or.Y][or.X] == ']' {
		p2.X -= 1
		act = true
	}

	if act {
		mv, cm := getTheMoves(mp, p, d, true)
		if !cm {
			return nil, false
		}
		mv2, cm := getTheMoves(mp, p2, d, true)
		if !cm {
			return nil, false
		}
		mv = MergeMaps(mv, mv2)
		mv[or] = struct{}{}
		return mv, true
	}

	return nil, false
}

func tryToMoveY(mp [][]byte, p common.Pos, d common.Pos) common.Pos {
	moves, canMove := getTheMoves(mp, p, d, false)
	if !canMove {
		return p
	}
	// do full movement
	posses := make([]common.Pos, len(moves))
	i := 0
	for p, _ := range moves {
		posses[i] = p
		i++
	}
	sort.Slice(posses, func(i, j int) bool {
		return posses[i].DirMul(d).Len() > posses[j].DirMul(d).Len()
	})

	// Handle all the pos changes
	for _, pp := range posses {
		tar := pp.Add(d)
		mp[tar.Y][tar.X] = mp[pp.Y][pp.X]
		mp[pp.Y][pp.X] = '.'
	}
	return p.Add(d)
}

func expand(r [][]byte) [][]byte {
	res := make([][]byte, 0)
	for _, ln := range r {
		lnn := make([]byte, 2*len(ln))
		for i, c := range ln {
			lnn[i*2] = '.'
			lnn[i*2+1] = '.'
			if c == 'O' {
				lnn[i*2] = '['
				lnn[i*2+1] = ']'
			}
			if c == '#' {
				lnn[i*2] = '#'
				lnn[i*2+1] = '#'
			}
			if c == '@' {
				lnn[i*2] = '@'
			}
		}
		res = append(res, lnn)

	}
	return res
}
