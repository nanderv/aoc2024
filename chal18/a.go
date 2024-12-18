package chal18

import (
	"bufio"
	"fmt"
	"github.com/nanderv/aoc2024/common"
	"io"
)

type node struct {
	common.Pos
	dist int
}

func HashNode(xS int) func(pos common.Pos) int {
	return func(n common.Pos) int {
		return n.Y*xS + n.X
	}
}

type Locater interface {
	Has(common.Pos) bool
}

func (n node) NB(bounds common.Pos, locater Locater) []common.Pos {
	nbs := []common.Pos{{X: 1}, {X: -1}, {Y: 1}, {Y: -1}}
	res := make([]common.Pos, 0)
	for _, nb := range nbs {
		o := n.Add(nb)
		if locater.Has(o) {
			continue
		}
		if o.X < 0 || o.Y < 0 || o.X >= bounds.X || o.Y >= bounds.Y {
			continue
		}
		res = append(res, nb)
	}
	return res
}

func (n node) Move(p common.Pos) node {
	m := n.Add(p)

	n.Pos = m
	n.dist += 1
	return n
}

func Afunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var x, y, reads int

	_, err := fmt.Sscanf(scanner.Text(), "Size: %d,%d,%d", &x, &y, &reads)
	if err != nil {
		fmt.Println(err)
	}
	size := common.Pos{X: x, Y: y}

	fallingMemory := make([]common.Pos, 0)
	for scanner.Scan() {
		_, err = fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		if err != nil {
			fmt.Println(err)
		}
		fallingMemory = append(fallingMemory, common.Pos{X: x, Y: y})
	}
	mp := common.NewMiMap(HashNode(size.X))

	ls := []node{{Pos: common.Pos{X: 0, Y: 0}, dist: 0}}
	fallCounter := 0
	for i := fallCounter; i < reads; i++ {
		mp.Set(fallingMemory[i])
	}
	return sim(size, ls, mp)

}
