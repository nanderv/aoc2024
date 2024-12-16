package chal16

import (
	"bufio"
	"github.com/nanderv/aoc2024/common"
	"io"
	"sort"
)

type robot struct {
	p     common.Pos
	d     int
	score int
	prev  *robot
}

func (r robot) Hash() int {
	return r.p.X*40000 + r.p.Y*4 + r.d
}

type MiMap[T any] struct {
	m map[int]T
	f func(T) int
}

func (m *MiMap[T]) Len() int {
	return len(m.m)
}
func (m *MiMap[T]) Has(a T) bool {
	i := m.f(a)
	_, ok := m.m[i]
	return ok
}

func (m *MiMap[T]) Get(a T) T {
	i := m.f(a)

	return m.m[i]
}
func (m *MiMap[T]) Set(a T) {
	i := m.f(a)
	m.m[i] = a
	return
}

func NewMiMap[T any](f func(T) int) MiMap[T] {
	return MiMap[T]{m: make(map[int]T), f: f}
}
func (r robot) Pos() common.Pos {
	return r.p
}

func (r robot) PosHash() int {
	return r.p.X*40000 + r.p.Y*4
}

// new robot, allowed to move there, end
func (r robot) Fwd(mp [][]byte) (robot, bool, bool) {
	res := robot{
		p:     r.p,
		d:     r.d,
		score: r.score,
		prev:  &r,
	}
	d := []common.Pos{{X: 1}, {Y: 1}, {X: -1}, {Y: -1}}
	pp := r.p.Add(d[r.d])
	if mp[pp.Y][pp.X] == '#' {
		return res, false, false
	}
	res.p = pp
	res.score += 1
	if mp[pp.Y][pp.X] == 'E' {
		return res, true, true
	}

	return res, true, false
}

func (r robot) Right(mp [][]byte) robot {
	res := robot{
		p:     r.p,
		d:     r.d,
		score: r.score,
		prev:  &r,
	}

	res.d = (r.d + 1) % 4
	res.score += 1000
	return res
}

func (r robot) Left(mp [][]byte) robot {
	res := robot{
		p:     r.p,
		d:     r.d,
		score: r.score,
		prev:  &r,
	}
	res.d = (r.d + 3) % 4
	res.score += 1000
	return res
}

func Afunc(file io.Reader) int {
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
	return i

}

func getResult(p common.Pos, mp [][]byte) (int, bool) {
	found := NewMiMap(robot.Hash)
	nexts := []robot{{p: p}}
	for len(nexts) > 0 {
		nxt := nexts[0]

		nexts = nexts[1:]

		nx, ok, done := nxt.Fwd(mp)
		if done {
			return nx.score, true
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
	return 0, false
}
