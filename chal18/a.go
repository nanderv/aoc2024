package chal18

import (
	"bufio"
	"fmt"
	"github.com/nanderv/aoc2024/common"
	"io"
)

type Node struct {
	common.Pos
	dist int
}

func HashNoded(xS int) func(n Node) int {
	return func(n Node) int {
		return n.Y*xS + n.X
	}
}
func HashNode(xS int) func(pos common.Pos) int {
	return func(n common.Pos) int {
		return n.Y*xS + n.X
	}
}

type Locater interface {
	Has(Node) bool
	Get(Node) Node
}

type Wrapped struct {
	l common.MiMap[Node]
	i int
}

func (w Wrapped) Has(n Node) bool {
	if !w.l.Has(n) {
		return false
	}
	return w.l.Get(n).dist <= w.i
}
func (w Wrapped) Get(n Node) Node {
	if w.Has(n) {
		return w.l.Get(n)
	}
	panic("OOPS")
}
func (n Node) NB(bounds common.Pos, locater Locater) []common.Pos {
	nbs := []common.Pos{{X: 1}, {X: -1}, {Y: 1}, {Y: -1}}
	res := make([]common.Pos, 0)
	for _, nb := range nbs {
		o := n.Add(nb)
		if locater.Has(Node{Pos: o}) {
			continue
		}
		if o.X < 0 || o.Y < 0 || o.X >= bounds.X || o.Y >= bounds.Y {
			continue
		}
		res = append(res, nb)
	}
	return res
}

func (n Node) Move(p common.Pos) Node {
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
	mp := common.NewMiMap(HashNoded(size.X))

	ls := []Node{{Pos: common.Pos{X: 0, Y: 0}, dist: 0}}
	fallCounter := 0
	for i := fallCounter; i < len(fallingMemory); i++ {
		mp.Set(Node{fallingMemory[i], i})
	}
	return sim(size, ls, mp, reads-1)

}
