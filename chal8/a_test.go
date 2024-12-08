package chal8_test

import (
	"github.com/nanderv/aoc2024/chal8"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChal1A(t *testing.T) {
	fl, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal8.Afunc(fl)
	assert.Equal(t, 14, res)
}

func TestChal1x(t *testing.T) {
	fl, err := os.Open("smol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal8.Afunc(fl)
	assert.Equal(t, 2, res)
}
