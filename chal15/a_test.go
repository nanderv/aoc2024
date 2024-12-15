package chal15_test

import (
	"github.com/nanderv/aoc2024/chal15"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChalExample(t *testing.T) {
	fl, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal15.Afunc(fl)
	assert.Equal(t, 10092, res)
}

func TestChalASmol(t *testing.T) {
	fl, err := os.Open("smol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal15.Afunc(fl)
	assert.Equal(t, 2028, res)
}
