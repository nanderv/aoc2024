package chal13

import (
	"bufio"
	"fmt"
	"io"
)

type pos struct {
	x int
	y int
}

func (p pos) add(q pos) pos {
	return pos{p.x + q.x, p.y + q.y}
}

func (p pos) mul(n int) pos {
	return pos{p.x * n, p.y * n}
}
func (p pos) eq(q pos) bool {
	return p.x == q.x && p.y == q.y
}
func Afunc(file io.Reader) int {
	res := 0
	scanner := bufio.NewScanner(file)
	for {
		scanner.Scan()
		butALine := scanner.Text()
		butA := pos{}
		butB := pos{}
		fmt.Sscanf(butALine, "Button A: X+%d, Y+%d", &butA.x, &butA.y)

		scanner.Scan()
		butBLine := scanner.Text()
		fmt.Sscanf(butBLine, "Button B: X+%d, Y+%d", &butB.x, &butB.y)

		scanner.Scan()
		prizeLine := scanner.Text()
		prize := pos{}
		fmt.Sscanf(prizeLine, "Prize: X=%d, Y=%d", &prize.x, &prize.y)
		// Empty line

		pRes := 10000000

		found := false
		for a := 0; a < 110; a++ {
			for b := 0; b < 110; b++ {
				aa := butA.mul(a)
				bb := butB.mul(b)
				if aa.add(bb).eq(prize) {
					if a+3*b < pRes {

						pRes = 3*a + b
						found = true
					}
				}
			}
		}
		if found {
			res += pRes
		}
		if !scanner.Scan() {
			break
		}
	}
	return res
}
