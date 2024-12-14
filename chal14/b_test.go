package chal14_test

import (
	"github.com/nanderv/aoc2024/chal14"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChaAs(t *testing.T) {
	fl, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal14.Bfunc(fl)
	assert.Equal(t, 7687, res)
}
