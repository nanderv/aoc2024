package chal17

import (
	//"fmt"
	"github.com/nanderv/aoc2024/chal11"
	"io"
	"strconv"
	"strings"
)

func Bfunc(file io.Reader) any {
	regA, regB, regC, progIn := CreateProgInput(file)
	strProg := strings.Split(progIn, ",")
	prog := make([]int, 0)
	for _, pp := range strProg {
		prog = append(prog, chal11.Must(strconv.Atoi(pp)))
	}
	digitsFound := -1

	for i := 0; true; i++ {
		if i == regA {
			continue
		}
		out, _ := execProgram(progIn, i, regB, regC)
		nFound := -1

		for j := 0; j < len(out); j++ {
			if prog[len(prog)-1-j] != out[len(out)-1-j] {
				break
			}
			nFound = j
		}
		if nFound > digitsFound {
			digitsFound = nFound
			if digitsFound == len(prog)-1 {
				return i
			}
			i = i*8 - 8
			if i < 0 {
				i = 0
			}
		}
	}
	return 0
}
