package chal2_test

import (
	"github.com/nanderv/aoc2024/chal2"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestChal1B(t *testing.T) {
	str :=
		`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	res := chal2.Bfunc(strings.NewReader(str))
	assert.Equal(t, 4, res)
}
