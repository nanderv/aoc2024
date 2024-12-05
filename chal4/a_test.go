package chal4_test

import (
	"github.com/nanderv/aoc2024/chal4"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestChal1A(t *testing.T) {
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
	res := chal4.Afunc(strings.NewReader(str))
	assert.Equal(t, 18, res)
}
