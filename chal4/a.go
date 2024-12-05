package chal4

import (
	"bufio"
	"io"
)

func Afunc(file io.Reader) int {
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
			y: 0,
		},
		{
			x: -1,
			y: 0,
		},
		{
			y: 1,
			x: 0,
		},
		{
			y: -1,
			x: 0,
		},
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
			for _, d := range directions {
				x := i
				y := j
				found := true
				for _, l := range "XMAS" {
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
					res += 1
				}
			}
		}
	}
	return res
}
