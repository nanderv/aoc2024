package chal6_test

import (
	"github.com/nanderv/aoc2024/chal6"
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
	res := chal6.Bfunc(fl)
	assert.Equal(t, 1602, res)
}
