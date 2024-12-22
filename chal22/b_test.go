package chal22_test

import (
	"fmt"
	"github.com/nanderv/aoc2024/chal22"
	"github.com/nanderv/aoc2024/common"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestChalBMicro(t *testing.T) {
	str := `6A`

	v := chal22.Bfunc(strings.NewReader(str))
	assert.Equal(t, 126384, v)
}

func TestChalBItest(t *testing.T) {
	str := `029A`

	kp, s := chal22.GetBaseKeypad()
	rc1, s2 := chal22.GetRemoteKeypad()
	rc2, s3 := chal22.GetRemoteKeypad()
	kp.PrevDo = rc1.Press
	rc1.PrevDo = rc2.Press
	rc2.PrevDo = chal22.Last
	states := []common.Pos{s, s2, s3, common.Pos{}}
	fmt.Println(chal22.In(str, kp, states))
	//assert.Equal(t, 126384, v)
}
