package chal15

import (
	"bufio"
	"fmt"
	"github.com/nanderv/aoc2024/common"
	"io"
)

type robot struct {
	p common.Pos
}

func Afunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	mp := make([][]byte, 0)
	p := common.Pos{}
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == '@' {
				p.X = i
				p.Y = len(mp)
			}
		}
		if len(line) == 0 {
			break
		}
		mp = append(mp, []byte(line))
	}

	for scanner.Scan() {
		ln := scanner.Text()
		for _, c := range ln {
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
			//printMP(mp)
			//fmt.Println("")
		}
	}

	res := 0
	for y, ln := range mp {
		for x, c := range ln {
			if c == 'O' {
				res += y*100 + x
			}
		}
	}
	return res
}
func printMP(mp [][]byte) {
	for _, ln := range mp {
		fmt.Println(string(ln))
	}
}
func tryToMove(mp [][]byte, p common.Pos, d common.Pos) common.Pos {
	or := p.Add(d)

	if mp[or.Y][or.X] == '.' {
		return or
	}
	q := or
	for mp[q.Y][q.X] == 'O' || mp[q.Y][q.X] == '[' || mp[q.Y][q.X] == ']' {
		q = q.Add(d)
	}
	if mp[q.Y][q.X] == '#' {
		return p
	}
	mp[or.Y][or.X] = '.'
	mp[q.Y][q.X] = 'O'

	return or
}
