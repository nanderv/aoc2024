package chal3

import (
	"bytes"
	"fmt"
	"io"
)

func Bfunc(file io.Reader) any {
	txt, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	res := 0
	do := true
	for i := 0; i < len(txt); i++ {
		tzt := txt[i:]
		var a, b int
		_, err = fmt.Fscanf(bytes.NewReader(tzt), "mul(%d,%d)", &a, &b)
		if err == nil {
			if do {
				res = res + a*b
			}
		}
		_, err = fmt.Fscanf(bytes.NewReader(tzt), "do()")
		if err == nil {
			do = true
		}
		_, err = fmt.Fscanf(bytes.NewReader(tzt), "don't()")
		if err == nil {
			do = false
		}
	}
	return res
}
