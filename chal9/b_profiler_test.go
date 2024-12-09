package chal9_test

import (
	"github.com/nanderv/aoc2024/chal9"
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
	res := chal9.Bfunc(fl)
	assert.Equal(t, 249943041417600, res)
}
