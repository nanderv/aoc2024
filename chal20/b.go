package chal20

import (
	"io"
)

func Bfunc(file io.Reader) any {
	return Algo(file, 100, 20)
}
