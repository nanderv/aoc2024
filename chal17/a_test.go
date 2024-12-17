package chal17_test

import (
	"github.com/nanderv/aoc2024/chal17"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestChalExample(t *testing.T) {
	fl, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal17.Afunc(fl)
	assert.Equal(t, 7036, res)
}

func TestChalMini(t *testing.T) {
	str :=
		`Register A: 0
Register B: 0
Register C: 9

Program: 2,6
`

	res := chal17.Afunc(strings.NewReader(str))
	assert.Equal(t, 7036, res)
}
