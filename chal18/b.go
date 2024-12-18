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

	mp := common.NewMiMap(HashNode(size.X))

	ls := []node{{Pos: common.Pos{X: 0, Y: 0}, dist: 0}}

	for i := 0; i < len(fallingMemory); i++ {
		mp.Set(fallingMemory[i])
		r := sim(size, ls, mp)
		if r == -1 {
			return fmt.Sprintf("%d,%d", fallingMemory[i].X, fallingMemory[i].Y)
		}
	}
	panic("Impossible")
}

func sim(size common.Pos, ls []node, mp common.MiMap[common.Pos]) int {
	visited := common.NewMiMap(HashNode(size.X))

	for len(ls) > 0 {

		l := ls[0]
		ls = ls[1:]
		if l.X == size.X-1 && l.Y == size.Y-1 {
			return l.dist
		}
		nbs := l.NB(size, &mp)
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
