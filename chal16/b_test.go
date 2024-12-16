package chal16_test

import (
	"github.com/nanderv/aoc2024/chal15"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChalBSmol2(t *testing.T) {
	fl, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal15.Bfunc(fl)
	assert.Equal(t, 45, res)
}

func TestChalBlarger(t *testing.T) {
	fl, err := os.Open("larger.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal15.Bfunc(fl)
	assert.Equal(t, 64, res)
}
