package chal1_test

import (
	"github.com/nanderv/aoc2024/chal1"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestChal1A(t *testing.T) {
	str :=
		`3   4
4   3
2   5
1   3
3   9
3   3
`
	res := chal1.Afunc(strings.NewReader(str))
	assert.Equal(t, 11, res)
}
