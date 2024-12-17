package chal17_test

import (
	"github.com/nanderv/aoc2024/chal17"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChalBSmol2(t *testing.T) {
	fl, err := os.Open("ex2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal17.Bfunc(fl)
	assert.Equal(t, 117440, res)
}
