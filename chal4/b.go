package chal4

import (
	"bufio"
	"io"
)

func Bfunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	lines := make([][]byte, 0)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}
	directions := []struct {
		x int
		y int
	}{
		{
			x: 1,
			y: 1,
		},
		{
			x: -1,
			y: 1,
		},
		{
			x: 1,
			y: -1,
		},
		{
			x: -1,
			y: -1,
		},
	}
	res := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			lres := 0
			for _, d := range directions {
				x := i - d.x
				y := j - d.y
				found := true
				for _, l := range "MAS" {
					if x < 0 || x >= len(lines) {
						found = false
						break
					}
					if y < 0 || y >= len(lines[x]) {
						found = false
						break
					}
					if lines[x][y] != byte(l) {
						found = false
						break
					}
					x += d.x
					y += d.y
				}
				if found {
					lres += 1
				}
			}
			if lres == 2 {
				res += 1
			}
		}
	}
	return res
}
