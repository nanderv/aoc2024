package chal22

import (
	"bufio"
	"github.com/nanderv/aoc2024/common"
	"io"
	"strconv"
)

type dt struct {
	in []int
	v  int
}

func hashdt(i dt) int {
	res := 0
	for _, d := range i.in {
		res += d
		res *= 100
	}
	return res
}
func Bfunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	ims := common.NewMiMap(hashdt)
	for scanner.Scan() {
		line := scanner.Text()
		sec, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		btz := make([]int, 4)
		sec, btz[0] = Nxter(sec)
		sec, btz[1] = Nxter(sec)
		sec, btz[2] = Nxter(sec)
		rnd := common.NewMiMap(hashdt)
		for i := 0; i < 2000-3; i++ {
			sec, btz[3] = Nxter(sec)
			v := sec % 10
			va := dt{in: btz, v: v}
			if !rnd.Has(va) {
				if x, ok := ims.MGet(va); ok {
					va.v += x.v
				}
				ims.Set(va)
				rnd.Set(va)
			}
			btz[0], btz[1], btz[2] = btz[1], btz[2], btz[3]
		}
	}

	res := 0
	for _, v := range ims.Keys() {
		if v.v > res {
			res = v.v
		}
	}
	return res
}

// Nxter
func Nxter(prev int) (int, int) {
	nxt := Nxt(prev)
	return nxt, (nxt % 10) - (prev % 10)
}
