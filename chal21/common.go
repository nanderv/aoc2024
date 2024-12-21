package chal21

import (
	"github.com/nanderv/aoc2024/common"
)

type key struct {
	x int
	y int
	b byte
}

func (k key) hash() int {
	return k.y*1000 + k.x
}
func GetBaseKeypad() common.MiMap[key] {
	mp := common.NewMiMap(key.hash)
	mp.Set(key{1, 3, '1'})
	mp.Set(key{2, 3, '2'})
	mp.Set(key{3, 3, '3'})
	mp.Set(key{1, 2, '4'})
	mp.Set(key{2, 2, '5'})
	mp.Set(key{3, 2, '6'})
	mp.Set(key{1, 1, '7'})
	mp.Set(key{2, 1, '8'})
	mp.Set(key{3, 1, '9'})
	mp.Set(key{2, 4, '0'})
	mp.Set(key{3, 4, 'A'})
	return mp
}
func GetRemoteKeypad() common.MiMap[key] {
	mp := common.NewMiMap(key.hash)

	mp.Set(key{2, 1, '^'})
	mp.Set(key{3, 1, 'A'})
	mp.Set(key{1, 2, '<'})
	mp.Set(key{2, 2, 'v'})
	mp.Set(key{3, 2, '>'})
	return mp

}
