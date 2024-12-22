package chal22_test

import (
	"github.com/nanderv/aoc2024/chal22"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChalB(t *testing.T) {
	fl, err := os.Open("exampleB.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()

	v := chal22.Bfunc(fl)
	assert.Equal(t, 23, v)
}

func TestChalBFull(t *testing.T) {
	fl, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()

	v := chal22.Bfunc(fl)
	assert.Equal(t, 2018, v)
}
