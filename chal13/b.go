package chal13

import (
	"bufio"
	"fmt"
	"io"
)

func Bfunc(file io.Reader) any {
	res := 0
	scanner := bufio.NewScanner(file)
	for {
		scanner.Scan()
		butALine := scanner.Text()
		pA := pos{}
		pB := pos{}
		fmt.Sscanf(butALine, "Button A: X+%d, Y+%d", &pA.x, &pA.y)

		scanner.Scan()
		butBLine := scanner.Text()
		fmt.Sscanf(butBLine, "Button B: X+%d, Y+%d", &pB.x, &pB.y)

		scanner.Scan()
		prizeLine := scanner.Text()
		pP := pos{}
		fmt.Sscanf(prizeLine, "Prize: X=%d, Y=%d", &pP.x, &pP.y)
		pP = pP.add(pos{10000000000000, 10000000000000})
		// Empty line

		im := pB.y*pA.x - pA.y*pB.x
		if im == 0 {
			fmt.Println("Skipping due to zero issue")
			if !scanner.Scan() {
				break
			}
			continue
		}
		bIm := pP.y*pA.x - pP.x*pA.y
		b := bIm / im

		a := (pP.x - b*pB.x) / pA.x
		aa, bb := pA.mul(a), pB.mul(b)
		zz := aa.add(bb)
		if zz.eq(pP) {
			res += a*3 + b
		}
		if !scanner.Scan() {
			break
		}
	}
	return res
}
