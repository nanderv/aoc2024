package chal4_test

import (
	"github.com/nanderv/aoc2024/chal4"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestChal1B(t *testing.T) {
	str :=
		`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	res := chal4.Bfunc(strings.NewReader(str))
	assert.Equal(t, 9, res)
}
