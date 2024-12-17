package chal17

import (
	"bufio"
	"fmt"
	"github.com/nanderv/aoc2024/chal11"
	"github.com/nanderv/aoc2024/common"
	"io"
	"math"
	"strconv"
	"strings"
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
	regA, regB, regC, progIn := CreateProgInput(file)
	out, res := execProgram(progIn, regA, regB, regC)
	fmt.Println(out)
	fmt.Println(res)
	return 0
}

func CreateProgInput(file io.Reader) (int, int, int, string) {
	scanner := bufio.NewScanner(file)

	regA := 0
	regB := 0
	regC := 0
	scanner.Scan()
	line := scanner.Text()
	_, err := fmt.Sscanf(line, "Register A: %d", &regA)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	line = scanner.Text()
	_, err = fmt.Sscanf(line, "Register B: %d", &regB)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	line = scanner.Text()
	_, err = fmt.Sscanf(line, "Register C: %d", &regC)
	if err != nil {
		panic(err)
	}
	scanner.Scan()

	scanner.Scan()
	line = scanner.Text()

	var progIn string
	_, err = fmt.Sscanf(line, "Program: %s", &progIn)
	if err != nil {
		panic(err)
	}
	return regA, regB, regC, progIn
}

func execProgram(progIn string, regA int, regB int, regC int) ([]int, string) {
	strProg := strings.Split(progIn, ",")
	prog := make([]int, 0)
	for _, pp := range strProg {
		prog = append(prog, chal11.Must(strconv.Atoi(pp)))
	}
	pc := 0
	out := make([]int, 0)
	for pc < len(prog) {
		lit := prog[pc+1]
		combo := prog[pc+1]
		if lit == 4 {
			combo = regA
		}
		if lit == 5 {
			combo = regB
		}
		if lit == 6 {
			combo = regC
		}
		if lit == 7 {
			combo = -4
		}
		switch prog[pc] {
		case 0:
			if combo == -4 {
				panic("HI")
			}
			pp := int(math.Pow(2, float64(combo)))
			regA = regA / pp
			pc += 2
		case 1:
			regB = regB ^ lit
			pc += 2
		case 2:
			if combo == -4 {
				panic("HI")
			}
			regB = combo % 8
			pc += 2
		case 3:
			if regA == 0 {
				pc += 2
			} else {
				pc = lit
			}
		case 4:
			regB = regB ^ regC
			pc += 2
		case 5:
			if combo == -4 {
				panic("HI")
			}
			out = append(out, combo%8)
			pc += 2
		case 6:
			if combo == -4 {
				panic("HI")
			}
			regB = regA / int(math.Pow(2, float64(combo)))

			pc += 2
		case 7:
			if combo == -4 {
				panic("HI")
			}
			regC = regA / int(math.Pow(2, float64(combo)))
			pc += 2
		default:
			panic("invalid cmd")
		}
	}
	res := ""
	start := false
	for _, i := range out {
		if start {
			res += ","
		}
		res += strconv.Itoa(i)

		start = true
	}
	return out, res
}
