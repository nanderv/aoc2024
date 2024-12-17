package chal3

import (
	"bytes"
	"fmt"
	"io"
)

func Afunc(file io.Reader) any {
	txt, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	res := 0
	for i := 0; i < len(txt); i++ {
		tzt := txt[i:]
		var a, b int
		_, err = fmt.Fscanf(bytes.NewReader(tzt), "mul(%d,%d)", &a, &b)
		if err == nil {
			res = res + a*b
		}
	}
	return res
}
