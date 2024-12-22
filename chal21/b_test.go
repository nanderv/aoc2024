package chal21_test

import (
	"fmt"
	"github.com/nanderv/aoc2024/chal21"
	"github.com/nanderv/aoc2024/common"
	"testing"
)

func TestChalBItest(t *testing.T) {
	str := `029A`

	kp, s := chal21.GetBaseKeypad()
	rc1, s2 := chal21.GetRemoteKeypad()
	rc2, s3 := chal21.GetRemoteKeypad()
	kp.PrevDo = rc1.Press
	rc1.PrevDo = rc2.Press
	rc2.PrevDo = chal21.Last
	states := []common.Pos{s, s2, s3, common.Pos{}}
	fmt.Println(chal21.In(str, kp, states))
	//assert.Equal(t, 126384, v)
}
