package chal15_test

import (
	"github.com/nanderv/aoc2024/chal15"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChaAs(t *testing.T) {
	fl, err := os.Open("larger.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal15.Bfunc(fl)
	assert.Equal(t, 9021, res)
}

func TestChalBSmol2(t *testing.T) {
	fl, err := os.Open("smol2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal15.Bfunc(fl)
	assert.Equal(t, 618, res)
}

func TestChalBSmol(t *testing.T) {
	fl, err := os.Open("smol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal15.Bfunc(fl)
	assert.Equal(t, 1751, res)
}

func TestChalBExample(t *testing.T) {
	fl, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal15.Bfunc(fl)
	assert.Equal(t, 9021, res)
}
