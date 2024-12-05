package chal3_test

import (
	"github.com/nanderv/aoc2024/chal3"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestChal1B(t *testing.T) {
	str :=
		`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	res := chal3.Bfunc(strings.NewReader(str))
	assert.Equal(t, 48, res)
}
