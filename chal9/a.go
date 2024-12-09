package chal9

import (
	"bufio"
	"io"
)

type numType int

func cInt(c byte) numType {
	switch c {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	panic("OOF")
}

const fileMode = true
const emptyMode = false

func emptyGen(nums []numType) func() int {
	ptr := -1
	return func() int {
		for {
			ptr++
			if nums[ptr] == numType(-1) {
				return ptr
			}
		}
	}
}
func fullGen(startPt int, nums []numType) func() (int, bool) {
	ptr := startPt + 2
	return func() (int, bool) {
		for {
			ptr--
			if ptr < 0 {
				return 0, false
			}
			if nums[ptr] != numType(-1) {
				return ptr, true
			}
		}
	}
}
func Afunc(file io.Reader) int {
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	txt := []byte(scanner.Text())
	arr := make([]numType, len(txt)*10)
	for k, _ := range arr {
		arr[k] = -1
	}

	ptr := numType(0)
	mode := fileMode
	ctr := numType(0)
	for _, a := range txt {
		v := cInt(a)

		if mode == fileMode {
			for i := numType(0); i < v; i++ {
				arr[ptr+i] = ctr
			}
			ctr++
			mode = emptyMode
		} else if mode == emptyMode {
			mode = fileMode
		}
		ptr += v
	}
	fg := fullGen(int(ptr+5), arr)

	eg := emptyGen(arr)
	for {
		pF, ok := fg()
		if !ok {
			break
		}
		pE := eg()
		if pE > pF {
			break
		}
		arr[pE] = arr[pF]

		arr[pF] = -1
	}
	res := 0
	for k, v := range arr {
		if int(v) == 0 {
		}
		if int(v) > 0 {
			res += k * int(v)
		}
	}
	return res
}
