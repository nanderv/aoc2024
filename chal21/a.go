package chal21

import (
	"bufio"
	"fmt"
	"github.com/nanderv/aoc2024/common"
	"io"
)

func Afunc(file io.Reader) any {
	scanner := bufio.NewScanner(file)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		x := getSq(GetBaseKeypad(), line)
		fmt.Println(x)
		x = getSq(GetRemoteKeypad(), x)
		fmt.Println(x)
		x = getSq(GetRemoteKeypad(), x)
		fmt.Println(x, len(x))
		res += len(x)
	}

	return res
}

func getSq(keypad common.MiMap[key], keys string) string {

	res := ""
	var start key
	for _, ky := range keypad.Keys() {
		if ky.b == 'A' {
			start = ky
		}
	}
	for _, k := range keys {
		s, nextSt := navTo(keypad, start, byte(k))
		res = res + s + "A"
		start = nextSt
	}

	return res
}
func navTo(keypad common.MiMap[key], start key, target byte) (string, key) {
	resultMap := common.NewMiMap(key.hash)
	dirs := []key{{x: 1, y: 0, b: '>'}, {x: -1, y: 0, b: '<'}, {x: 0, y: -1, b: '^'}, {x: 0, y: 1, b: 'v'}}
	rtz := []key{start}
	resPos := key{}
	resStr := ""
	for len(rtz) > 0 {

		nxt := rtz[0]
		rtz = rtz[1:]
		for _, d := range dirs {
			tg := key{x: d.x + nxt.x, y: d.y + nxt.y, b: d.b}
			if keypad.Has(tg) && !resultMap.Has(tg) {
				resultMap.Set(tg)
				kk := keypad.Get(tg)
				if kk.b == target {
					resPos = kk
					break
				}
				rtz = append(rtz, tg)
			}
		}
	}
	rollback := resPos
	for rollback.x != start.x || rollback.y != start.y {
		rm := resultMap.Get(rollback)
		if rm.b == '^' {
			resStr = "^" + resStr
			rollback.y += 1
		} else if rm.b == '>' {
			resStr = ">" + resStr
			rollback.x -= 1
		} else if rm.b == '<' {
			resStr = "<" + resStr
			rollback.x += 1
		} else if rm.b == 'v' {
			resStr = "v" + resStr
			rollback.y -= 1
		} else {
			panic("impossible, rm")
		}
	}
	return resStr, resPos
	// todo: analyse result map
}
