package chal10_test

import (
	"github.com/nanderv/aoc2024/chal10"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChal1BProfile(t *testing.T) {
	fl, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal10.Bfunc(fl)
	assert.Equal(t, 6420913943576, res)
}
