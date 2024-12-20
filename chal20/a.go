package chal20

import (
	"io"
)

func Afunc(file io.Reader) any {
	return Algo(file, 100, 2)
}
