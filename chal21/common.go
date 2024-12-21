package chal21

import (
	"fmt"
	"github.com/nanderv/aoc2024/common"
	"strconv"
	"strings"
)

type key struct {
	common.Pos
	b byte
}

func (k key) hash() int {
	return k.Y*100 + k.X
}

type KeyPad struct {
	KeyMap common.MiMap[key]
	DoAct  func(p Metapos, machineNr int, c byte) (Metapos, bool)

	PrevDo func(PressIn) PressOut // we know that we start from 'A'
}

// In starts the reverse approach: start from the last machine and work backwards.
func In(s string, kp KeyPad, states []common.Pos) int {
	res := 0
	for _, c := range s {
		po := kp.Press(PressIn{states, byte(c)})
		res += po.res
		states[0] = po.p
	}
	return res
}
func Last(_ PressIn) PressOut {
	return PressOut{common.Pos{}, 1}
}

type PressIn struct {
	p []common.Pos
	b byte
}

func (p PressIn) Hash() string {
	res := ""
	for _, pp := range p.p {
		res += "+" + strconv.Itoa(pHash(pp))
	}

	return res + "Z" + string(p.b)
}

type PressOut struct {
	p   common.Pos
	res int
}

func (k *KeyPad) Press(ii PressIn) PressOut {
	p := ii.p
	b := ii.b
	pz := p[0]
	var ky key
	for _, kyy := range k.KeyMap.Keys() {
		if kyy.b == b {
			ky = kyy
		}
	}
	diff := ky.Pos.Sub(pz)
	yp := ""
	xp := ""
	if diff.Y < 0 {
		// up then left or right
		yp = strings.Repeat("^", -diff.Y)
	}
	if diff.Y > 0 {
		yp = strings.Repeat("v", diff.Y)
	}
	if diff.X < 0 {
		xp = strings.Repeat("<", -diff.X)
	}
	if diff.X > 0 {
		xp = strings.Repeat(">", diff.X)
	}

	newList := p[1:]
	opt1Lst := make([]common.Pos, len(newList))
	opt2Lst := make([]common.Pos, len(newList))
	copy(opt1Lst, newList)
	copy(opt2Lst, newList)
	res1 := 0
	has1 := false
	if k.KeyMap.Has(key{Pos: pz.Add(common.Pos{Y: diff.Y})}) {
		has1 = true
		for _, c := range yp + xp + "A" {
			pi := k.PrevDo(PressIn{opt1Lst, byte(c)})
			res1 += pi.res
			opt1Lst[0] = pi.p
		}
	}
	res2 := 0
	has2 := false
	if k.KeyMap.Has(key{Pos: pz.Add(common.Pos{X: diff.X})}) {
		has2 = true
		for _, c := range xp + yp + "A" {
			pi := k.PrevDo(PressIn{opt2Lst, byte(c)})
			res2 += pi.res
			opt2Lst[0] = pi.p
		}
	}
	if !has2 && !has1 {
		fmt.Println("STart of oof")
		fmt.Println(pz, ky)
		fmt.Println(pz.Add(common.Pos{X: diff.X}))
		fmt.Println(pz.Add(common.Pos{Y: diff.Y}))
		fmt.Println(p)
		panic("OOF")
	}

	if !has1 {
		copy(newList, opt2Lst)
		return PressOut{ky.Pos, res2}
	}
	if !has2 {
		copy(newList, opt1Lst)
		return PressOut{ky.Pos, res1}
	}
	if res1 < res2 {
		copy(newList, opt1Lst)

		return PressOut{ky.Pos, res1}
	}
	copy(newList, opt2Lst)
	return PressOut{ky.Pos, res2}
}

// CalcMove

// Out and RcAct together form the forward simulation approach: this essentially performs BFS on the whole thing
func Out(s string) func(p Metapos, machineNr int, c byte) (Metapos, bool) {
	return func(p Metapos, machineNr int, c byte) (Metapos, bool) {
		if c != s[p.corChars] {
			return p, false
		}
		p.corChars++
		return p, true
	}
}

// Takes a slice of position data (one per machine), how far into the result we are, and the next input character,
// Returns two booleans:
// Whether the move was valid
// Whether we should increment the result counter
func (k *KeyPad) RcAct(p Metapos, machineNr int, c byte) (Metapos, bool) {
	pos := p.pad[machineNr]
	pad := make([]common.Pos, len(p.pad))
	for i, pp := range p.pad {
		pad[i] = pp
	}

	dirs := map[byte]common.Pos{
		'^': {Y: -1},
		'v': {Y: 1},
		'<': {X: -1},
		'>': {X: 1},
	}
	if d, ok := dirs[c]; ok {
		np := pos.Add(d)

		if k.KeyMap.Has(key{Pos: np}) {
			pad[machineNr] = np
		}
		// do directions
	}
	if c == 'A' {
		// Do action on next machine
		return k.DoAct(p, machineNr+1, k.KeyMap.Get(key{Pos: pos}).b)

	}

	return Metapos{
		pad:      pad,
		corChars: p.corChars,
		dist:     p.dist,
		inp:      p.inp,
	}, true
}

func GetBaseKeypad() (KeyPad, common.Pos) {
	mp := common.NewMiMap(key.hash)
	mp.Set(key{common.Pos{X: 1, Y: 3}, '1'})
	mp.Set(key{common.Pos{X: 2, Y: 3}, '2'})
	mp.Set(key{common.Pos{X: 3, Y: 3}, '3'})
	mp.Set(key{common.Pos{X: 1, Y: 2}, '4'})
	mp.Set(key{common.Pos{X: 2, Y: 2}, '5'})
	mp.Set(key{common.Pos{X: 3, Y: 2}, '6'})
	mp.Set(key{common.Pos{X: 1, Y: 1}, '7'})
	mp.Set(key{common.Pos{X: 2, Y: 1}, '8'})
	mp.Set(key{common.Pos{X: 3, Y: 1}, '9'})
	mp.Set(key{common.Pos{X: 2, Y: 4}, '0'})
	mp.Set(key{common.Pos{X: 3, Y: 4}, 'A'})
	return KeyPad{KeyMap: mp}, common.Pos{3, 4}
}
func GetRemoteKeypad() (KeyPad, common.Pos) {
	mp := common.NewMiMap(key.hash)

	mp.Set(key{Pos: common.Pos{X: 2, Y: 1}, b: '^'})
	mp.Set(key{Pos: common.Pos{X: 3, Y: 1}, b: 'A'})
	mp.Set(key{Pos: common.Pos{X: 1, Y: 2}, b: '<'})
	mp.Set(key{Pos: common.Pos{X: 2, Y: 2}, b: 'v'})
	mp.Set(key{Pos: common.Pos{X: 3, Y: 2}, b: '>'})
	return KeyPad{KeyMap: mp}, common.Pos{X: 3, Y: 1}
}

type Metapos struct {
	pad      []common.Pos
	corChars int
	dist     int
	inp      string
}

func pHash(p common.Pos) int {
	return p.X*100 + p.Y
}
func (m Metapos) Hash() int {
	res := 0
	for _, p := range m.pad {
		res += pHash(p)
		res *= 1000
	}
	res *= 100
	//res += m.dist
	res *= 1000
	res += m.corChars
	return res
}

func (m Metapos) SlowHash() string {
	res := ""
	for _, p := range m.pad {
		res += ":" + strconv.Itoa(pHash(p))
	}
	res += strconv.Itoa(m.corChars)
	return res
}
