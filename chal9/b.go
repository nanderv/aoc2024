package chal9

import (
	"bufio"
	"io"
)

type slot struct {
	ptr  int
	size int
}

func getEmptySlot(nums []numType, size int) slot {
	ptr := -1

	for {
		for {
			ptr++
			if nums[ptr] == numType(-1) {
				break
			}
		}

		sz := 0
		lp := ptr
		for nums[ptr] == numType(-1) && sz < 12 {
			sz++
			ptr++
		}
		if sz >= size {
			return slot{ptr: lp, size: sz}
		}
	}
}

func getFiles(startPt int, nums []numType) func() (slot, bool) {
	ptr := startPt + 2
	lowest := numType(99999999)
	return func() (slot, bool) {
		for {
			ptr--
			if ptr < 0 {
				return slot{}, false
			}
			if nums[ptr] != numType(-1) && nums[ptr] < lowest {
				v := nums[ptr]
				lowest = v
				size := 0
				for ptr >= 0 && nums[ptr] == v {
					size++
					ptr--
				}
				ptr++

				return slot{ptr, size}, true
			}
		}
	}
}
func Bfunc(file io.Reader) any {
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
	fg := getFiles(int(ptr+5), arr)

	for {
		pF, ok := fg()
		if !ok {
			break
		}
		eF := getEmptySlot(arr, pF.size)
		if pF.ptr < eF.ptr {
			continue
		}
		for i := 0; i < pF.size; i++ {
			arr[eF.ptr+i] = arr[pF.ptr+i]
			arr[pF.ptr+i] = -1
		}

	}
	res := 0
	for k, v := range arr {
		if int(v) > 0 {
			res += k * int(v)
		}
	}
	return res
}
