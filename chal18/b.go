package chal18

import (
	"bufio"
	"fmt"
	"github.com/nanderv/aoc2024/common"
	"io"
	"sort"
)

func Bfunc(file io.Reader) any {
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
	for i := 0; i < len(fallingMemory); i++ {
		mp.Set(Node{Pos: fallingMemory[i], dist: i + 1})
	}
	low := 0
	high := len(fallingMemory) - 1
	for high-low > 1 {

		mid := (low + high) / 2

		ls := []Node{{Pos: common.Pos{X: 0, Y: 0}, dist: 0}}

		r := sim(size, ls, mp, mid)
		if r == -1 {
			if low == mid {
				return fmt.Sprintf("%d,%d", fallingMemory[low].X, fallingMemory[low].Y)
			}
			high = mid
		} else {
			low = mid
		}
	}
	return fmt.Sprintf("%d,%d", fallingMemory[low].X, fallingMemory[low].Y)

}

func sim(size common.Pos, ls []Node, mp common.MiMap[Node], dist int) int {
	visited := common.NewMiMap(HashNode(size.X))
	wr := Wrapped{l: mp, i: dist}
	for len(ls) > 0 {

		l := ls[0]
		ls = ls[1:]
		if l.X == size.X-1 && l.Y == size.Y-1 {
			return l.dist
		}
		nbs := l.NB(size, &wr)
		for _, nb := range nbs {
			m := l.Move(nb)
			if !visited.Has(m.Pos) {
				ls = append(ls, l.Move(nb))
			}
			visited.Set(m.Pos)
		}
		sort.Slice(ls, func(i, j int) bool { return ls[i].dist < ls[j].dist })
	}
	return -1
}
