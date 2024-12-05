package chal3_test

import (
	"github.com/nanderv/aoc2024/chal3"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestChal1A(t *testing.T) {
	str :=
		`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	res := chal3.Afunc(strings.NewReader(str))
	assert.Equal(t, 161, res)
}
