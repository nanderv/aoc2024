package chal17

import (
	"fmt"
	"io"
)

func Bfunc(file io.Reader) int {
	regA, regB, regC, progIn := CreateProgInput(file)
	regAOrig := regA
	for i := 0; i < 1000000000; i++ {
		if i%100000 == 0 {
			fmt.Println(i)
		}
		if i == regAOrig {
			continue
		}

		_, res := execProgram(progIn, i, regB, regC)
		//fmt.Println(res)
		//fmt.Println(progIn)
		//fmt.Println()
		if res == progIn {
			return i
		}
	}
	return 0
}
